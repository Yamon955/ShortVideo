// Code generated by trpc-go/trpc-cmdline v1.0.9. DO NOT EDIT.
// source: hello.proto

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

// HelloWorldServiceService defines service.
type HelloWorldServiceService interface {
	// Hello Hello says hello.
	Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error)
}

func HelloWorldServiceService_Hello_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &HelloRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(HelloWorldServiceService).Hello(ctx, reqbody.(*HelloRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// HelloWorldServiceServer_ServiceDesc descriptor for server.RegisterService.
var HelloWorldServiceServer_ServiceDesc = server.ServiceDesc{
	ServiceName: "trpc.shortvideo.test.HelloWorldService",
	HandlerType: ((*HelloWorldServiceService)(nil)),
	Methods: []server.Method{
		{
			Name: "/trpc.shortvideo.test.HelloWorldService/Hello",
			Func: HelloWorldServiceService_Hello_Handler,
		},
	},
}

// RegisterHelloWorldServiceService registers service.
func RegisterHelloWorldServiceService(s server.Service, svr HelloWorldServiceService) {
	if err := s.Register(&HelloWorldServiceServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("HelloWorldService register error:%v", err))
	}
}

// START --------------------------------- Default Unimplemented Server Service --------------------------------- START

type UnimplementedHelloWorldService struct{}

// Hello Hello says hello.
func (s *UnimplementedHelloWorldService) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return nil, errors.New("rpc Hello of service HelloWorldService is not implemented")
}

// END --------------------------------- Default Unimplemented Server Service --------------------------------- END

// END ======================================= Server Service Definition ======================================= END

// START ======================================= Client Service Definition ======================================= START

// HelloWorldServiceClientProxy defines service client proxy
type HelloWorldServiceClientProxy interface {
	// Hello Hello says hello.
	Hello(ctx context.Context, req *HelloRequest, opts ...client.Option) (rsp *HelloResponse, err error)
}

type HelloWorldServiceClientProxyImpl struct {
	client client.Client
	opts   []client.Option
}

var NewHelloWorldServiceClientProxy = func(opts ...client.Option) HelloWorldServiceClientProxy {
	return &HelloWorldServiceClientProxyImpl{client: client.DefaultClient, opts: opts}
}

func (c *HelloWorldServiceClientProxyImpl) Hello(ctx context.Context, req *HelloRequest, opts ...client.Option) (*HelloResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/trpc.shortvideo.test.HelloWorldService/Hello")
	msg.WithCalleeServiceName(HelloWorldServiceServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("shortvideo")
	msg.WithCalleeServer("test")
	msg.WithCalleeService("HelloWorldService")
	msg.WithCalleeMethod("Hello")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &HelloResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

// END ======================================= Client Service Definition ======================================= END
