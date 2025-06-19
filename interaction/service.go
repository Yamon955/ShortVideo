package interaction

import (
	"context"
	"github.com/Yamon955/ShortVideo/protocol/interaction/pb"
)

type interactionSvrImpl struct {
}

func newInteractionSvr() *interactionSvrImpl {
	return &interactionSvrImpl{}
}

func (s *interactionSvrImpl) Like(ctx context.Context, req *pb.LikeReq) (*pb.LikeRsp, error) {

	return &pb.LikeRsp{}, nil
}

func (s *interactionSvrImpl) Follow(ctx context.Context, req *pb.FollowReq) (*pb.FollowRsp, error) {

	return &pb.FollowRsp{}, nil
}

func (s *interactionSvrImpl) Commit(ctx context.Context, req *pb.CommitReq) (*pb.CommitRsp, error) {
	return &pb.CommitRsp{}, nil
}

func (s *interactionSvrImpl) Report(ctx context.Context, req *pb.ReportReq) (*pb.ReportRsp, error) {
	return &pb.ReportRsp{}, nil
}
