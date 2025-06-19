package logic

import (
	"context"
	"github.com/Yamon955/ShortVideo/interaction/entity/errcode"
	"github.com/Yamon955/ShortVideo/protocol/interaction/pb"
	"trpc.group/trpc-go/trpc-go/errs"
)

func (h *handlerImpl) HandleCommit(ctx context.Context, req *pb.CommitReq) (*pb.CommitRsp, error) {
	if req.GetOperationType() == pb.OPERATION_TYPES_DO {
		return commit(ctx, req)
	} else if req.GetOperationType() == pb.OPERATION_TYPES_UNDO {
		return undoCommit(ctx, req)
	}
	return nil, errs.New(errcode.ErrInvalidOperationType, "非法操作类型")
}

// 评论
func commit(ctx context.Context, req *pb.CommitReq) (*pb.CommitRsp, error) {
	return &pb.CommitRsp{}, nil
}

// 取消评论
func undoCommit(ctx context.Context, req *pb.CommitReq) (*pb.CommitRsp, error) {
	return &pb.CommitRsp{}, nil
}
