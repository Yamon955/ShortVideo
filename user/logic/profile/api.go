// Package profile 个人资料相关包
package profile

import (
	"context"

	"github.com/Yamon955/ShortVideo/user/pb"
	"github.com/Yamon955/ShortVideo/user/repo/mysql"
)

// Handler 用户资料处理器
type Handler interface {
	// HandleBatchGetProfile 处理用户资料获取
	HandleBatchGetProfile(ctx context.Context, req *pb.BatchGetProfileReq, rsp *pb.BatchGetProfileRsp) error
	// HandleSetProfile 处理用户资料修改
	HandleSetProfile(ctx context.Context, req *pb.SetProfileReq, rsp *pb.SetProfileRsp) error
}

// NewHandler 新建一个用户资料处理器
func NewHandler() Handler {
	return &handlerImpl{
		db: mysql.NewDBClient(),
	}
}
