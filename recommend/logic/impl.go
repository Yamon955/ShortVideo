package logic

import (
	"context"
	"strconv"

	"github.com/Yamon955/ShortVideo/protocol/recommend/pb"
	"github.com/Yamon955/ShortVideo/recommend/entity/def"
	"github.com/redis/go-redis/v9"
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
		bfKey := strconv.FormatUint(req.Uid, 10) + def.BloomFilter1Keysuffix
		vidsAfterFilter := make([]string, 0, 10)
		for _, vid := range vids {
			exist := h.rdb.BFExists(ctx, bfKey, vid).Val()
			if !exist {
				vidsAfterFilter = append(vidsAfterFilter, vid)
			}
		}
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
