// Code generated by trpc-go/trpc-cmdline v1.0.9. DO NOT EDIT.
// source: user.proto

package pb

import (
	"context"
	"errors"
	"fmt"

	_ "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/codec"
	_ "trpc.group/trpc-go/trpc-go/http"
	"trpc.group/trpc-go/trpc-go/server"
)

// START ======================================= Server Service Definition ======================================= START

// UserService defines service.
type UserService interface {
	Register(ctx context.Context, req *RegisterReq) (*RegisterRsp, error)

	Login(ctx context.Context, req *LoginReq) (*LoginRsp, error)

	BatchGetProfile(ctx context.Context, req *BatchGetProfileReq) (*BatchGetProfileRsp, error)

	SetProfile(ctx context.Context, req *SetProfileReq) (*SetProfileRsp, error)
}

func UserService_Register_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &RegisterReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(UserService).Register(ctx, reqbody.(*RegisterReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func UserService_Login_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &LoginReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(UserService).Login(ctx, reqbody.(*LoginReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func UserService_BatchGetProfile_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &BatchGetProfileReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(UserService).BatchGetProfile(ctx, reqbody.(*BatchGetProfileReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func UserService_SetProfile_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &SetProfileReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(UserService).SetProfile(ctx, reqbody.(*SetProfileReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// UserServer_ServiceDesc descriptor for server.RegisterService.
var UserServer_ServiceDesc = server.ServiceDesc{
	ServiceName: "trpc.shortvideo.user.User",
	HandlerType: ((*UserService)(nil)),
	Methods: []server.Method{
		{
			Name: "/trpc.shortvideo.user.User/Register",
			Func: UserService_Register_Handler,
		},
		{
			Name: "/trpc.shortvideo.user.User/Login",
			Func: UserService_Login_Handler,
		},
		{
			Name: "/trpc.shortvideo.user.User/BatchGetProfile",
			Func: UserService_BatchGetProfile_Handler,
		},
		{
			Name: "/trpc.shortvideo.user.User/SetProfile",
			Func: UserService_SetProfile_Handler,
		},
	},
}

// RegisterUserService registers service.
func RegisterUserService(s server.Service, svr UserService) {
	if err := s.Register(&UserServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("User register error:%v", err))
	}
}

// START --------------------------------- Default Unimplemented Server Service --------------------------------- START

type UnimplementedUser struct{}

func (s *UnimplementedUser) Register(ctx context.Context, req *RegisterReq) (*RegisterRsp, error) {
	return nil, errors.New("rpc Register of service User is not implemented")
}
func (s *UnimplementedUser) Login(ctx context.Context, req *LoginReq) (*LoginRsp, error) {
	return nil, errors.New("rpc Login of service User is not implemented")
}
func (s *UnimplementedUser) BatchGetProfile(ctx context.Context, req *BatchGetProfileReq) (*BatchGetProfileRsp, error) {
	return nil, errors.New("rpc BatchGetProfile of service User is not implemented")
}
func (s *UnimplementedUser) SetProfile(ctx context.Context, req *SetProfileReq) (*SetProfileRsp, error) {
	return nil, errors.New("rpc SetProfile of service User is not implemented")
}

// END --------------------------------- Default Unimplemented Server Service --------------------------------- END

// END ======================================= Server Service Definition ======================================= END

// START ======================================= Client Service Definition ======================================= START

// UserClientProxy defines service client proxy
type UserClientProxy interface {
	Register(ctx context.Context, req *RegisterReq, opts ...client.Option) (rsp *RegisterRsp, err error)

	Login(ctx context.Context, req *LoginReq, opts ...client.Option) (rsp *LoginRsp, err error)

	BatchGetProfile(ctx context.Context, req *BatchGetProfileReq, opts ...client.Option) (rsp *BatchGetProfileRsp, err error)

	SetProfile(ctx context.Context, req *SetProfileReq, opts ...client.Option) (rsp *SetProfileRsp, err error)
}

type UserClientProxyImpl struct {
	client client.Client
	opts   []client.Option
}

var NewUserClientProxy = func(opts ...client.Option) UserClientProxy {
	return &UserClientProxyImpl{client: client.DefaultClient, opts: opts}
}

func (c *UserClientProxyImpl) Register(ctx context.Context, req *RegisterReq, opts ...client.Option) (*RegisterRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/trpc.shortvideo.user.User/Register")
	msg.WithCalleeServiceName(UserServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("shortvideo")
	msg.WithCalleeServer("user")
	msg.WithCalleeService("User")
	msg.WithCalleeMethod("Register")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &RegisterRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *UserClientProxyImpl) Login(ctx context.Context, req *LoginReq, opts ...client.Option) (*LoginRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/trpc.shortvideo.user.User/Login")
	msg.WithCalleeServiceName(UserServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("shortvideo")
	msg.WithCalleeServer("user")
	msg.WithCalleeService("User")
	msg.WithCalleeMethod("Login")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &LoginRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *UserClientProxyImpl) BatchGetProfile(ctx context.Context, req *BatchGetProfileReq, opts ...client.Option) (*BatchGetProfileRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/trpc.shortvideo.user.User/BatchGetProfile")
	msg.WithCalleeServiceName(UserServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("shortvideo")
	msg.WithCalleeServer("user")
	msg.WithCalleeService("User")
	msg.WithCalleeMethod("BatchGetProfile")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &BatchGetProfileRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *UserClientProxyImpl) SetProfile(ctx context.Context, req *SetProfileReq, opts ...client.Option) (*SetProfileRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/trpc.shortvideo.user.User/SetProfile")
	msg.WithCalleeServiceName(UserServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("shortvideo")
	msg.WithCalleeServer("user")
	msg.WithCalleeService("User")
	msg.WithCalleeMethod("SetProfile")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &SetProfileRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

// END ======================================= Client Service Definition ======================================= END
