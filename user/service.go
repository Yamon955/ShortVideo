package main

import (
	"context"

	"github.com/Yamon955/ShortVideo/user/logic/auth"
	"github.com/Yamon955/ShortVideo/user/pb"
	"trpc.group/trpc-go/trpc-go/errs"
)

type userSvrImpl struct{
	authHandler auth.AuthHandler
}

func newUserSvr() *userSvrImpl {
	return &userSvrImpl{
		authHandler: auth.NewAuthHandler(),
	}
}

// Register 注册功能
func (s *userSvrImpl) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterRsp, error) {
	rsp := &pb.RegisterRsp{}
	err := s.authHandler.HandleRegister(ctx, req)
	if err != nil {
		formatRegisterRsp(rsp, int64(errs.Code(err)), err.Error())
		return rsp, err
	}
	formatRegisterRsp(rsp, 0, "注册成功，赶快去登录吧～")
	return rsp, nil
}

// Login 登录功能
func (s *userSvrImpl) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRsp, error) {

	return nil, nil
}

func formatRegisterRsp(rsp *pb.RegisterRsp, code int64, msg string) {
	rsp.Code = code
	rsp.Msg  = msg
}

func formatLoginRsp(rsp *pb.LoginRsp, code int64, msg string) {
	rsp.Code = code
	rsp.Msg  = msg
}