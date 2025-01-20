package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/textproto"
	"strings"
	"time"

	"github.com/Yamon955/ShortVideo/gateway/entity"
	"github.com/google/uuid"
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

// 当调用的服务返回 err != nil 时，rsp.Body 会被清空，因此，需要将 err 信息转成 json 形式放到 http.rsp 中
type errRsp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

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
	// 后端服务可能用到的 metadata
	opts = append(opts, getOpts(ctx, r)...)
	rspBody := &codec.Body{}
	log.InfoContextf(ctx, "URL:%s forward success", r.URL.Path)
	if err := client.DefaultClient.Invoke(ctx, reqBody, rspBody, opts...); err != nil {
		// err != nil 时，rsp.Body 会被清空
		rsp := &errRsp{
			Code: int(errs.Code(err)),
			Msg:  errs.Msg(err),
		}
		rspBody.Data, _ = json.Marshal(rsp)
	}
	if err := transTRPCRsp(w, r, rspBody); err != nil {
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
		log.Errorf("GetSerializationType failed, err:%v", err)
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

// transTRPCRsp trpc.rsp -> http.rsp
func transTRPCRsp(w http.ResponseWriter, _ *http.Request, rspBody *codec.Body) error {
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

// getOpts 后端服务可能用到的 metadata
func getOpts(ctx context.Context, r *http.Request) (opts []client.Option) {
	loginUID := ctx.Value("sv_login_uid")
	var uid string
	if loginUID != nil {
		uid = loginUID.(string)
	}
	tranceID := getTraceID(ctx, uid)
	log.ErrorContextf(ctx, "uid[%s]", uid)
	opts = []client.Option{
		client.WithMetaData("sv_login_uid", []byte(uid)),
		client.WithMetaData("sv_trace_id", []byte(tranceID)),
	}
	return
}

// getLoginUID 从 cookie 中获取 uid
func getLoginUID(r *http.Request) string {
	if r == nil {
		return ""
	}
	loginUID, err := r.Cookie("sv_login_uid")
	if err != nil {
		return ""
	}
	return loginUID.Value
}

func getTraceID(_ context.Context, uin string) string {
	if uin == "" {
		uuidWithHyphen := uuid.New()
		uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
		return fmt.Sprintf("%s_%d_%d", uuid, time.Now().UnixMilli(), rand.Uint32())
	}
	return fmt.Sprintf("%s_%d_%d", uin, time.Now().UnixMilli(), rand.Uint32())
}
