package auth

import (
	"context"

	"github.com/Yamon955/ShortVideo/user/pb"
)

// NewAuthHandler 创建一个用户注册、登录的处理器
func NewAuthHandler() AuthHandler {
	return &authHandlerandlerImpl{}
}

// AuthHandler 用户注册、登录处理接口
type AuthHandler interface {
	// HandleRegister 用户注册
	HandleRegister(ctx context.Context, req *pb.RegisterReq) error
	// HandleLogin 用户登录
	HandleLogin(ctx context.Context, req *pb.LoginReq, rsp *pb.LoginRsp) error
}