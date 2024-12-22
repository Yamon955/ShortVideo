package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/textproto"
	"time"

	"github.com/Yamon955/ShortVideo/gateway/entity"
	ctrpc "trpc.group/trpc-go/trpc-gateway/common/trpc"
	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/codec"
	"trpc.group/trpc-go/trpc-go/errs"
	"trpc.group/trpc-go/trpc-go/log"
)

var (
	// 通过 textproto.CanonicalMIMEHeaderKey 转成 http.Header 键值标准形式
	canonicalContentType   = textproto.CanonicalMIMEHeaderKey("Content-Type")
	canonicalContentLength = textproto.CanonicalMIMEHeaderKey("Content-Length")
)

func forwardTRPC(w http.ResponseWriter, r *http.Request, routeConf *entity.RouteConf) error {
	timeout, err := time.ParseDuration(routeConf.Timeout)
	if err != nil {
		timeout = 2 * time.Second
	}
	ctx, msg := codec.WithCloneMessage(r.Context())
	defer codec.PutBackMessage(msg)
	reqBody, err := transTRPCReqBody(ctx, r)
	if err != nil {
		return err
	}
	calleeMethod := routeConf.CalleeMethod
	msg.WithCallerMethod(r.URL.Path)
	msg.WithClientRPCName(calleeMethod)
	opts := []client.Option{
		client.WithTimeout(timeout),
		client.WithProtocol(entity.ProtocolTRPC),
		client.WithServiceName(routeConf.Service),
		client.WithCalleeMethod(calleeMethod),
		client.WithTarget(routeConf.Target),
		client.WithCurrentSerializationType(codec.SerializationTypeNoop), // 透传即可，无需进行序列化
		client.WithCurrentCompressType(codec.CompressTypeNoop),           // 不要自动解压缩
	}
	rspBody := &codec.Body{}
	if err := client.DefaultClient.Invoke(ctx, reqBody, rspBody, opts...); err != nil{
		return err
	}
	if err := transHTTPRsp(w, r, rspBody); err != nil {
		return errs.New(entity.ErrTransHTTPRsp, "trpc.rsp trans to http.rsp err")
	}
	return nil

}

// transTRPCReqBody http.req.body -> trpc.req.body
func transTRPCReqBody(ctx context.Context, r *http.Request) (*codec.Body, error) {
	reqBody := &codec.Body{}
	msg := codec.Message(ctx)
	if r.Method == http.MethodGet {
		msg.WithSerializationType(codec.SerializationTypeGet)
		reqBody.Data = []byte(r.URL.RawQuery)
		return reqBody, nil
	}
	serializationType, err := ctrpc.GetSerializationType(r.Header.Get(canonicalContentType))
	if err != nil {
		log.Errorf("GetSerializationType failed, err:%v", serializationType, err)
		return nil, errs.New(entity.ErrContentType, "invalid content type")
	}
	msg.WithSerializationType(serializationType)
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf("read http.req.body failed, err:%v", err)
		return nil, errs.New(entity.ErrReadHTTPReqBody, "read HTTP req body err")
	}
	reqBody.Data = data
	return reqBody, nil
}

// transHTTPRsp trpc.rsp -> http.rsp
func transHTTPRsp(w http.ResponseWriter, _ *http.Request, rspBody *codec.Body) error {
	defaultContentType := "application/json"
	w.Header().Set(canonicalContentType, defaultContentType)
	w.Header().Set(canonicalContentLength, fmt.Sprintf("%d", len(rspBody.Data)))
	_, err := w.Write(rspBody.Data)
	if err != nil {
		log.Errorf("write to http.rsp.body failed, err:%v", err)
		return err
	}
	return nil
}
