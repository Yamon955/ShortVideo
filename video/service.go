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

func newVideoSvr() *videoSvrImpl {
	return &videoSvrImpl{
		handler: logic.NewHandler(),
	}
}

// Publish 视频发布
func (v *videoSvrImpl) Publish(ctx context.Context, req *pb.PublishReq) (*pb.PublishRsp, error) {
	rsp, err := v.handler.HandlePublish(ctx, req)
	if err != nil {
		log.ErrorContextf(ctx, "HandlePublish failed, err:%v", err)
		return nil, err
	}
	return rsp, nil
}

// GetFeeds 获取推荐视频列表
func (v *videoSvrImpl) GetFeeds(ctx context.Context, req *pb.GetFeedsReq) (*pb.GetFeedsRsp, error) {
	return &pb.GetFeedsRsp{}, nil
}

// GetPublishList 获取用户发布列表
func (v *videoSvrImpl) GetPublishList(ctx context.Context, req *pb.GetPublishListReq) (*pb.GetPublishListRsp, error) {
	rsp, err := v.handler.HandleGetPublishList(ctx, req)
	if err != nil {
		log.ErrorContextf(ctx, "HandleGetPublishList failed, err:%v", err)
		return nil, err
	}
	return rsp, nil
}

// GetLikeList 获取赞过列表
func (v *videoSvrImpl) GetLikeList(ctx context.Context, req *pb.GetLikeListReq) (*pb.GetLikeListRsp, error) {

	return &pb.GetLikeListRsp{}, nil
}

// GetCollectList 获取收藏列表
func (v *videoSvrImpl) GetCollectList(ctx context.Context, req *pb.GetCollectListReq) (*pb.GetCollectListRsp, error) {

	return &pb.GetCollectListRsp{}, nil
}
