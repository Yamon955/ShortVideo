package logic

import (
	"context"
	"github.com/Yamon955/ShortVideo/interaction/entity/errcode"
	"github.com/Yamon955/ShortVideo/protocol/interaction/pb"
	"trpc.group/trpc-go/trpc-go/errs"
)

func (h *handlerImpl) HandleFollow(ctx context.Context, req *pb.FollowReq) (*pb.FollowRsp, error) {
	if req.GetOperationType() == pb.OPERATION_TYPES_DO {
		return follow(ctx, req)
	} else if req.GetOperationType() == pb.OPERATION_TYPES_UNDO {
		return undoFollow(ctx, req)
	}
	return nil, errs.New(errcode.ErrInvalidOperationType, "非法操作类型")
}

// 关注
func follow(ctx context.Context, req *pb.FollowReq) (*pb.FollowRsp, error) {
	return &pb.FollowRsp{}, nil
}

// 取消关注
func undoFollow(ctx context.Context, req *pb.FollowReq) (*pb.FollowRsp, error) {
	return &pb.FollowRsp{}, nil
}
