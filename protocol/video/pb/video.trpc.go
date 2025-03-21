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

	GetPublishList(ctx context.Context, req *GetPublishListReq) (*GetPublishListRsp, error)

	GetLikeList(ctx context.Context, req *GetLikeListReq) (*GetLikeListRsp, error)

	GetCollectList(ctx context.Context, req *GetCollectListReq) (*GetCollectListRsp, error)

	WatchVideo(ctx context.Context, req *WatchVideoReq) (*WatchVideoRsp, error)
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

func VideoService_GetLikeList_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &GetLikeListReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(VideoService).GetLikeList(ctx, reqbody.(*GetLikeListReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func VideoService_GetCollectList_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &GetCollectListReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(VideoService).GetCollectList(ctx, reqbody.(*GetCollectListReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func VideoService_WatchVideo_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &WatchVideoReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(VideoService).WatchVideo(ctx, reqbody.(*WatchVideoReq))
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
			Name: "/trpc.shortvideo.video.Video/GetPublishList",
			Func: VideoService_GetPublishList_Handler,
		},
		{
			Name: "/trpc.shortvideo.video.Video/GetLikeList",
			Func: VideoService_GetLikeList_Handler,
		},
		{
			Name: "/trpc.shortvideo.video.Video/GetCollectList",
			Func: VideoService_GetCollectList_Handler,
		},
		{
			Name: "/trpc.shortvideo.video.Video/WatchVideo",
			Func: VideoService_WatchVideo_Handler,
		},
	},
}

// RegisterVideoService registers service.
func RegisterVideoService(s server.Service, svr VideoService) {
	if err := s.Register(&VideoServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("Video register error:%v", err))
	}
}

// PublishService defines service.
type PublishService interface {
	Publish(ctx context.Context, req *PublishReq) (*PublishRsp, error)
}

func PublishService_Publish_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &PublishReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(PublishService).Publish(ctx, reqbody.(*PublishReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// PublishServer_ServiceDesc descriptor for server.RegisterService.
var PublishServer_ServiceDesc = server.ServiceDesc{
	ServiceName: "trpc.shortvideo.video.Publish",
	HandlerType: ((*PublishService)(nil)),
	Methods: []server.Method{
		{
			Name: "/trpc.shortvideo.video.Publish/Publish",
			Func: PublishService_Publish_Handler,
		},
	},
}

// RegisterPublishService registers service.
func RegisterPublishService(s server.Service, svr PublishService) {
	if err := s.Register(&PublishServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("Publish register error:%v", err))
	}
}

// START --------------------------------- Default Unimplemented Server Service --------------------------------- START

type UnimplementedVideo struct{}

func (s *UnimplementedVideo) GetFeeds(ctx context.Context, req *GetFeedsReq) (*GetFeedsRsp, error) {
	return nil, errors.New("rpc GetFeeds of service Video is not implemented")
}
func (s *UnimplementedVideo) GetPublishList(ctx context.Context, req *GetPublishListReq) (*GetPublishListRsp, error) {
	return nil, errors.New("rpc GetPublishList of service Video is not implemented")
}
func (s *UnimplementedVideo) GetLikeList(ctx context.Context, req *GetLikeListReq) (*GetLikeListRsp, error) {
	return nil, errors.New("rpc GetLikeList of service Video is not implemented")
}
func (s *UnimplementedVideo) GetCollectList(ctx context.Context, req *GetCollectListReq) (*GetCollectListRsp, error) {
	return nil, errors.New("rpc GetCollectList of service Video is not implemented")
}
func (s *UnimplementedVideo) WatchVideo(ctx context.Context, req *WatchVideoReq) (*WatchVideoRsp, error) {
	return nil, errors.New("rpc WatchVideo of service Video is not implemented")
}

type UnimplementedPublish struct{}

func (s *UnimplementedPublish) Publish(ctx context.Context, req *PublishReq) (*PublishRsp, error) {
	return nil, errors.New("rpc Publish of service Publish is not implemented")
}

// END --------------------------------- Default Unimplemented Server Service --------------------------------- END

// END ======================================= Server Service Definition ======================================= END

// START ======================================= Client Service Definition ======================================= START

// VideoClientProxy defines service client proxy
type VideoClientProxy interface {
	GetFeeds(ctx context.Context, req *GetFeedsReq, opts ...client.Option) (rsp *GetFeedsRsp, err error)

	GetPublishList(ctx context.Context, req *GetPublishListReq, opts ...client.Option) (rsp *GetPublishListRsp, err error)

	GetLikeList(ctx context.Context, req *GetLikeListReq, opts ...client.Option) (rsp *GetLikeListRsp, err error)

	GetCollectList(ctx context.Context, req *GetCollectListReq, opts ...client.Option) (rsp *GetCollectListRsp, err error)

	WatchVideo(ctx context.Context, req *WatchVideoReq, opts ...client.Option) (rsp *WatchVideoRsp, err error)
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

func (c *VideoClientProxyImpl) GetLikeList(ctx context.Context, req *GetLikeListReq, opts ...client.Option) (*GetLikeListRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/trpc.shortvideo.video.Video/GetLikeList")
	msg.WithCalleeServiceName(VideoServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("shortvideo")
	msg.WithCalleeServer("video")
	msg.WithCalleeService("Video")
	msg.WithCalleeMethod("GetLikeList")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &GetLikeListRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *VideoClientProxyImpl) GetCollectList(ctx context.Context, req *GetCollectListReq, opts ...client.Option) (*GetCollectListRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/trpc.shortvideo.video.Video/GetCollectList")
	msg.WithCalleeServiceName(VideoServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("shortvideo")
	msg.WithCalleeServer("video")
	msg.WithCalleeService("Video")
	msg.WithCalleeMethod("GetCollectList")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &GetCollectListRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *VideoClientProxyImpl) WatchVideo(ctx context.Context, req *WatchVideoReq, opts ...client.Option) (*WatchVideoRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/trpc.shortvideo.video.Video/WatchVideo")
	msg.WithCalleeServiceName(VideoServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("shortvideo")
	msg.WithCalleeServer("video")
	msg.WithCalleeService("Video")
	msg.WithCalleeMethod("WatchVideo")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &WatchVideoRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

// PublishClientProxy defines service client proxy
type PublishClientProxy interface {
	Publish(ctx context.Context, req *PublishReq, opts ...client.Option) (rsp *PublishRsp, err error)
}

type PublishClientProxyImpl struct {
	client client.Client
	opts   []client.Option
}

var NewPublishClientProxy = func(opts ...client.Option) PublishClientProxy {
	return &PublishClientProxyImpl{client: client.DefaultClient, opts: opts}
}

func (c *PublishClientProxyImpl) Publish(ctx context.Context, req *PublishReq, opts ...client.Option) (*PublishRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/trpc.shortvideo.video.Publish/Publish")
	msg.WithCalleeServiceName(PublishServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("shortvideo")
	msg.WithCalleeServer("video")
	msg.WithCalleeService("Publish")
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

// END ======================================= Client Service Definition ======================================= END
