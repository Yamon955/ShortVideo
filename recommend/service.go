package main

import (
	"context"

	"github.com/Yamon955/ShortVideo/protocol/recommend/pb"
	"github.com/Yamon955/ShortVideo/recommend/logic"
)

type recommendSvrImpl struct {
	handler logic.Handler
}

func (s *recommendSvrImpl) FeedsRecommend(
	ctx context.Context,
	req *pb.FeedsRecommendReq,
) (*pb.FeedsRecommendRsp, error) {

	return &pb.FeedsRecommendRsp{}, nil
}

func newRecommendSvr() *recommendSvrImpl {
	return &recommendSvrImpl{
		handler: logic.NewHandler(),
	}
}
