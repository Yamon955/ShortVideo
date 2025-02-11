package logic

import (
	"context"

	"github.com/Yamon955/ShortVideo/protocol/recommend/pb"
	"github.com/Yamon955/ShortVideo/recommend/repo/redis"
)

// Handler 推荐服务处理器
type Handler interface {
	// HandleFeedsRecommend 获取推荐 feeds
	HandleFeedsRecommend(ctx context.Context, req *pb.FeedsRecommendReq) (*pb.FeedsRecommendRsp, error)
}

func NewHandler() Handler {
	return &handlerImpl{
		rdb: redis.GetClient(),
	}
}
