package logic

import (
	"context"
	"github.com/Yamon955/ShortVideo/interaction/entity/errcode"
	"github.com/Yamon955/ShortVideo/protocol/interaction/pb"
	"trpc.group/trpc-go/trpc-go/errs"
)

func (h *handlerImpl) HandleLike(ctx context.Context, req *pb.LikeReq) (*pb.LikeRsp, error) {
	if req.GetOperationType() == pb.OPERATION_TYPES_DO {
		return like(ctx, req)
	} else if req.GetOperationType() == pb.OPERATION_TYPES_UNDO {
		return undoLike(ctx, req)
	}
	return nil, errs.New(errcode.ErrInvalidOperationType, "非法操作类型")
}

// 点赞
func like(ctx context.Context, req *pb.LikeReq) (*pb.LikeRsp, error) {
	return &pb.LikeRsp{}, nil
}

// 取消点赞
func undoLike(ctx context.Context, req *pb.LikeReq) (*pb.LikeRsp, error) {
	return &pb.LikeRsp{}, nil
}
