package logic

import (
	"context"

	"github.com/Yamon955/ShortVideo/protocol/video/pb"
)

func (h *handlerImpl) HandleGetFeeds(ctx context.Context, req *pb.GetFeedsReq) (*pb.GetFeedsRsp, error) {
	return &pb.GetFeedsRsp{}, nil
}
