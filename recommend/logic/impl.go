package logic

import (
	"context"
	"strconv"

	"github.com/Yamon955/ShortVideo/protocol/recommend/pb"
	"github.com/Yamon955/ShortVideo/recommend/entity/def"
	"github.com/redis/go-redis/v9"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

type handlerImpl struct {
	rdb *redis.Client
}

// HandleFeedsRecommend 获取推荐 feeds
func (h *handlerImpl) HandleFeedsRecommend(
	ctx context.Context,
	req *pb.FeedsRecommendReq,
) (*pb.FeedsRecommendRsp, error) {
	if req.StartIndex < 0 {
		req.StartIndex = 0
	}
	recommendedVids := make([]uint64, 0, 2*def.RecommendFeesCount)
	startIndex := req.GetStartIndex()
	for len(recommendedVids) < def.RecommendFeesCount {
		// 查询 startIndex 开始的 20 个视频 id
		endIndex := startIndex + def.VideosCountEscape - 1
		vids := h.rdb.ZRevRange(ctx, def.VideosInTwoMonthRedisKey, startIndex, endIndex).Val()
		// 更新下一次查询的起始索引位置
		startIndex = endIndex + 1
		// 布隆过滤器去除用户已经看过的视频
		uid := strconv.FormatUint(req.Uid, 10)
		vidsAfterFilter := h.duplicateVidByBloomFilter(ctx, uid, vids)
		// 标签值匹配筛选用户
		for _, vid := range vidsAfterFilter {
			vidTagKey := def.VideoTagKeyPrefix + vid
			uidTagKey := def.UserTagKeyPrefix + strconv.FormatUint(req.GetUid(), 10)
			// 标签值匹配
			val := h.rdb.BitOpAnd(ctx, vidTagKey, uidTagKey).Val()
			if val >= 0 {
				recommendedVid, _ := strconv.ParseUint(vid, 10, 64)
				recommendedVids = append(recommendedVids, recommendedVid)
			}
		}
	}
	log.Infof("recommendedVids:%v", recommendedVids)
	return &pb.FeedsRecommendRsp{
		Vids:      recommendedVids,
		NextIndex: startIndex,
	}, nil
}

// duplicateVidByBloomFilter 利用两个布隆过滤器去重
func (h *handlerImpl) duplicateVidByBloomFilter(ctx context.Context, uid string, vids []string) []string {
	var handlers []func() error
	var exists1, exists2 []bool
	// 布隆过滤器去除用户已经看过的视频, 两个都要判断
	bfKey1 := uid + def.BloomFilter1Keysuffix
	bfKey2 := uid + def.BloomFilter2Keysuffix
	handlers = append(handlers, func() error {
		exists1 = h.rdb.BFMExists(ctx, bfKey1, vids).Val()
		return nil
	})
	handlers = append(handlers, func() error {
		exists2 = h.rdb.BFMExists(ctx, bfKey2, vids).Val()
		return nil
	})
	_ = trpc.GoAndWait(handlers...)

	vidsAfterFilter := make([]string, 0, len(vids))
	for i, vid := range vids {
		// 只要可能在其中一个里面，就表示用户已经刷到过，去重
		if exists1[i] || exists2[i] {
			continue
		}
		vidsAfterFilter = append(vidsAfterFilter, vid)
	}
	return vidsAfterFilter
}
