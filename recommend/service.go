package main

import (
	"context"

	"github.com/Yamon955/ShortVideo/protocol/recommend/pb"
	"github.com/Yamon955/ShortVideo/recommend/entity/errcode"
	"github.com/Yamon955/ShortVideo/recommend/logic"
	"trpc.group/trpc-go/trpc-go/errs"
	"trpc.group/trpc-go/trpc-go/log"
)

type recommendSvrImpl struct {
	handler logic.Handler
}

func (s *recommendSvrImpl) FeedsRecommend(
	ctx context.Context,
	req *pb.FeedsRecommendReq,
) (*pb.FeedsRecommendRsp, error) {
	log.ErrorContextf(ctx, "Recommend")
	rsp, err := s.handler.HandleFeedsRecommend(ctx, req)
	if err != nil {
		log.ErrorContextf(ctx, "HandleFeedsRecommend failed, err:%v", err)
		return nil, errs.New(errcode.ErrRecommendFeedsFailed, "Get recommend feeds failed")
	}
	return rsp, nil
}

func newRecommendSvr() *recommendSvrImpl {
	return &recommendSvrImpl{
		handler: logic.NewHandler(),
	}
}
