package profile

import (
	"context"

	"github.com/Yamon955/ShortVideo/protocol/user/pb"
	"github.com/Yamon955/ShortVideo/user/repo/mysql"
)

// handlerImpl 用户资料处理器实现
type handlerImpl struct {
	db mysql.DBClient
}

// HandleBatchGetProfile 处理用户资料获取
func (h *handlerImpl) HandleBatchGetProfile(
	ctx context.Context,
	req *pb.BatchGetProfileReq,
	rsp *pb.BatchGetProfileRsp) error {

	return nil
}

// HandleSetProfile 处理用户资料修改
func (h *handlerImpl) HandleSetProfile(ctx context.Context, req *pb.SetProfileReq, rsp *pb.SetProfileRsp) error {

	return nil
}
