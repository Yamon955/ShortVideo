package main

import (
	"context"

	"github.com/Yamon955/ShortVideo/protocol/video/pb"
	"github.com/Yamon955/ShortVideo/video/logic"
	"trpc.group/trpc-go/trpc-go/log"
)

type videoSvrImpl struct {
	handler logic.Handler
}

func (v *videoSvrImpl) GetFeeds(ctx context.Context, req *pb.GetFeedsReq) (*pb.GetFeedsRsp, error) {
	return &pb.GetFeedsRsp{}, nil
}

func (v *videoSvrImpl) Publish(ctx context.Context, req *pb.PublishReq) (*pb.PublishRsp, error) {
	rsp, err := v.handler.HandlePublish(ctx, req)
	if err != nil {
		log.ErrorContextf(ctx, "HandlePublish failed, err:%v", err)
		return nil, err
	}
	return rsp, nil
}

func (v *videoSvrImpl) GetPublishList(ctx context.Context, req *pb.GetPublishListReq) (*pb.GetPublishListRsp, error) {
	return &pb.GetPublishListRsp{}, nil
}
