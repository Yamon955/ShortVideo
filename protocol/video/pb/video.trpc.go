// Code generated by trpc-go/trpc-cmdline v1.0.9. DO NOT EDIT.
// source: video.proto

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

// VideoService defines service.
type VideoService interface {
	GetFeeds(ctx context.Context, req *GetFeedsReq) (*GetFeedsRsp, error)

	Publish(ctx context.Context, req *PublishReq) (*PublishRsp, error)

	GetPublishList(ctx context.Context, req *GetPublishListReq) (*GetPublishListRsp, error)
}

func VideoService_GetFeeds_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &GetFeedsReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(VideoService).GetFeeds(ctx, reqbody.(*GetFeedsReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func VideoService_Publish_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &PublishReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(VideoService).Publish(ctx, reqbody.(*PublishReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func VideoService_GetPublishList_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &GetPublishListReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(VideoService).GetPublishList(ctx, reqbody.(*GetPublishListReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// VideoServer_ServiceDesc descriptor for server.RegisterService.
var VideoServer_ServiceDesc = server.ServiceDesc{
	ServiceName: "trpc.shortvideo.video.Video",
	HandlerType: ((*VideoService)(nil)),
	Methods: []server.Method{
		{
			Name: "/trpc.shortvideo.video.Video/GetFeeds",
			Func: VideoService_GetFeeds_Handler,
		},
		{
			Name: "/trpc.shortvideo.video.Video/Publish",
			Func: VideoService_Publish_Handler,
		},
		{
			Name: "/trpc.shortvideo.video.Video/GetPublishList",
			Func: VideoService_GetPublishList_Handler,
		},
	},
}

// RegisterVideoService registers service.
func RegisterVideoService(s server.Service, svr VideoService) {
	if err := s.Register(&VideoServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("Video register error:%v", err))
	}
}

// START --------------------------------- Default Unimplemented Server Service --------------------------------- START

type UnimplementedVideo struct{}

func (s *UnimplementedVideo) GetFeeds(ctx context.Context, req *GetFeedsReq) (*GetFeedsRsp, error) {
	return nil, errors.New("rpc GetFeeds of service Video is not implemented")
}
func (s *UnimplementedVideo) Publish(ctx context.Context, req *PublishReq) (*PublishRsp, error) {
	return nil, errors.New("rpc Publish of service Video is not implemented")
}
func (s *UnimplementedVideo) GetPublishList(ctx context.Context, req *GetPublishListReq) (*GetPublishListRsp, error) {
	return nil, errors.New("rpc GetPublishList of service Video is not implemented")
}

// END --------------------------------- Default Unimplemented Server Service --------------------------------- END

// END ======================================= Server Service Definition ======================================= END

// START ======================================= Client Service Definition ======================================= START

// VideoClientProxy defines service client proxy
type VideoClientProxy interface {
	GetFeeds(ctx context.Context, req *GetFeedsReq, opts ...client.Option) (rsp *GetFeedsRsp, err error)

	Publish(ctx context.Context, req *PublishReq, opts ...client.Option) (rsp *PublishRsp, err error)

	GetPublishList(ctx context.Context, req *GetPublishListReq, opts ...client.Option) (rsp *GetPublishListRsp, err error)
}

type VideoClientProxyImpl struct {
	client client.Client
	opts   []client.Option
}

var NewVideoClientProxy = func(opts ...client.Option) VideoClientProxy {
	return &VideoClientProxyImpl{client: client.DefaultClient, opts: opts}
}

func (c *VideoClientProxyImpl) GetFeeds(ctx context.Context, req *GetFeedsReq, opts ...client.Option) (*GetFeedsRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/trpc.shortvideo.video.Video/GetFeeds")
	msg.WithCalleeServiceName(VideoServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("shortvideo")
	msg.WithCalleeServer("video")
	msg.WithCalleeService("Video")
	msg.WithCalleeMethod("GetFeeds")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &GetFeedsRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *VideoClientProxyImpl) Publish(ctx context.Context, req *PublishReq, opts ...client.Option) (*PublishRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/trpc.shortvideo.video.Video/Publish")
	msg.WithCalleeServiceName(VideoServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("shortvideo")
	msg.WithCalleeServer("video")
	msg.WithCalleeService("Video")
	msg.WithCalleeMethod("Publish")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &PublishRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *VideoClientProxyImpl) GetPublishList(ctx context.Context, req *GetPublishListReq, opts ...client.Option) (*GetPublishListRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/trpc.shortvideo.video.Video/GetPublishList")
	msg.WithCalleeServiceName(VideoServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("shortvideo")
	msg.WithCalleeServer("video")
	msg.WithCalleeService("Video")
	msg.WithCalleeMethod("GetPublishList")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &GetPublishListRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

// END ======================================= Client Service Definition ======================================= END