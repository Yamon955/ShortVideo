// Package auth 登录鉴权相关包
package auth

import (
	"context"

	"github.com/Yamon955/ShortVideo/protocol/user/pb"
	"github.com/Yamon955/ShortVideo/user/repo/mysql"
	"github.com/Yamon955/ShortVideo/user/repo/redis"
)

// Handler 用户注册、登录处理器
type Handler interface {
	// HandleRegister 用户注册
	HandleRegister(ctx context.Context, req *pb.RegisterReq) error
	// HandleLogin 用户登录
	HandleLogin(ctx context.Context, req *pb.LoginReq, rsp *pb.LoginRsp) error
}

// NewHandler 创建一个用户注册、登录的处理器
func NewHandler() Handler {
	return &handlerImpl{
		db:  mysql.NewDBClient(),
		rdb: redis.GetClient(),
	}
}
