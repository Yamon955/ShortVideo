package main

import (
	"context"

	"github.com/Yamon955/ShortVideo/protocol/user/pb"
	"github.com/Yamon955/ShortVideo/user/logic/auth"
	"github.com/Yamon955/ShortVideo/user/logic/profile"
)

type userSvrImpl struct {
	authHandler    auth.Handler
	profileHandler profile.Handler
}

func newUserSvr() *userSvrImpl {
	return &userSvrImpl{
		authHandler:    auth.NewHandler(),
		profileHandler: profile.NewHandler(),
	}
}

// Register 注册功能
func (s *userSvrImpl) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterRsp, error) {
	rsp := &pb.RegisterRsp{}
	err := s.authHandler.HandleRegister(ctx, req)
	if err != nil {
		return nil, err
	}
	formatRsp(rsp, 0, "注册成功，赶快去登录吧～")
	return rsp, nil
}

// Login 登录功能
func (s *userSvrImpl) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRsp, error) {
	rsp := &pb.LoginRsp{}
	if err := s.authHandler.HandleLogin(ctx, req, rsp); err != nil {
		return nil, err
	}
	formatRsp(rsp, 0, "登录成功")
	return rsp, nil
}

// BatchGetProfile 用户资料获取，支持批量方式
func (s *userSvrImpl) BatchGetProfile(ctx context.Context, req *pb.BatchGetProfileReq) (*pb.BatchGetProfileRsp, error) {
	rsp := &pb.BatchGetProfileRsp{}
	if err := s.profileHandler.HandleBatchGetProfile(ctx, req, rsp); err != nil {
		return rsp, nil
	}
	return rsp, nil
}

// SetProfile 用户资料修改
func (s *userSvrImpl) SetProfile(ctx context.Context, req *pb.SetProfileReq) (*pb.SetProfileRsp, error) {
	rsp := &pb.SetProfileRsp{}
	if err := s.profileHandler.HandleSetProfile(ctx, req, rsp); err != nil {
		return rsp, nil
	}
	return rsp, nil
}

// formatRsp 填充回包的 code、msg 字段
func formatRsp(rsp interface{}, code int64, msg string) {
	switch rsp.(type) {
	case *pb.RegisterRsp:
		tmp := rsp.(*pb.RegisterRsp)
		tmp.Code = code
		tmp.Msg = msg
	case *pb.LoginRsp:
		tmp := rsp.(*pb.LoginRsp)
		tmp.Code = code
		tmp.Msg = msg
	}
}
