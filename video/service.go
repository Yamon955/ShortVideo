package main

import (
	"context"

	"github.com/Yamon955/ShortVideo/protocol/video/pb"
)

type videoSvrImpl struct {
}

func (v *videoSvrImpl) GetFeeds(ctx context.Context, req *pb.GetFeedsReq) (*pb.GetFeedsRsp, error) {
	return &pb.GetFeedsRsp{}, nil
}

func (v *videoSvrImpl) Publish(ctx context.Context, req *pb.PublishReq) (*pb.PublishRsp, error) {

	return &pb.PublishRsp{}, nil
}

func (v *videoSvrImpl) GetPublishList(ctx context.Context, req *pb.GetPublishListReq) (*pb.GetPublishListRsp, error) {
	return &pb.GetPublishListRsp{}, nil
}
