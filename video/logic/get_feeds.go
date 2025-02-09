package logic

import (
	"context"

	"github.com/Yamon955/ShortVideo/protocol/video/pb"
)

// HandleGetFeeds 获取推荐视频列表
func (h *handlerImpl) HandleGetFeeds(ctx context.Context, req *pb.GetFeedsReq) (*pb.GetFeedsRsp, error) {
	return &pb.GetFeedsRsp{}, nil
}
