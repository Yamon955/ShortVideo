package logic

import (
	"context"

	"github.com/Yamon955/ShortVideo/protocol/video/pb"
)

func (h *handlerImpl) HandleGetPublishList(ctx context.Context, req *pb.GetPublishListReq) (*pb.GetPublishListRsp, error) {
	return &pb.GetPublishListRsp{}, nil
}
