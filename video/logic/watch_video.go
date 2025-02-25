package logic

import (
	"context"
	"strings"

	"github.com/Yamon955/ShortVideo/protocol/video/pb"
	"github.com/Yamon955/ShortVideo/video/entity/def"
	"trpc.group/trpc-go/trpc-go"
)

// HandleWatchVideo 处理用户观看的 video
func (h *handlerImpl) HandleWatchVideo(ctx context.Context, req *pb.WatchVideoReq) (*pb.WatchVideoRsp, error) {
	uid := string(trpc.GetMetaData(ctx, "sv_login_uid"))
	vid := req.GetVid()

	// 当前正在使用的布隆过滤器的 key
	currentBFKey := h.rdb.Get(ctx, def.UserCurrentBloomFilterPrefix+uid).Val()

	// vid 添加到当前的布隆过滤器中
	addSuccess := h.rdb.BFAdd(ctx, currentBFKey, vid).Val()
	if addSuccess {
		return &pb.WatchVideoRsp{}, nil
	}
	// 添加失败，当前的布隆过滤器满了
	var oldBFKey string
	if strings.HasSuffix(currentBFKey, "1") {
		oldBFKey = uid + def.BloomFilter2Keysuffix
	} else {
		oldBFKey = uid + def.BloomFilter1Keysuffix
	}
	// 清理旧的布隆过滤器 删除->重建
	_ = h.rdb.Del(ctx, oldBFKey)
	_ = h.rdb.BFReserve(ctx, oldBFKey, def.BloomFilterErrRate, def.BloomFilterCapacity)
	// 重新指向新建的布隆过滤器
	_ = h.rdb.Set(ctx, def.UserCurrentBloomFilterPrefix+uid, oldBFKey, -1)
	// vid 加入清理之后的布隆过滤器中
	_ = h.rdb.BFAdd(ctx, oldBFKey, vid).Val()
	return &pb.WatchVideoRsp{}, nil
}
