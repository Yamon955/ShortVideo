package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Yamon955/ShortVideo/gateway/entity"
	ctrpc "trpc.group/trpc-go/trpc-gateway/common/trpc"
	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/codec"
	"trpc.group/trpc-go/trpc-go/errs"
	thttp "trpc.group/trpc-go/trpc-go/http"
	"trpc.group/trpc-go/trpc-go/log"
)

func forwardHTTP(w http.ResponseWriter, r *http.Request, routerConf *entity.RouteConf) error {
	rspBody := &codec.Body{}
	timeout, err := time.ParseDuration(routerConf.Timeout)
	if err != nil {
		timeout = 2 * time.Second
	}
	ctx := r.Context()
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	reqHead, reqBody, err := transHTTPReq(r)
	if err != nil {
		return err
	}
	method := routerConf.Method
	rpcName := method
	msg.WithClientRPCName(rpcName)
	serializationType, err := ctrpc.GetSerializationType(r.Header.Get(canonicalContentType))
	msg.WithSerializationType(serializationType)
	rsphead := &thttp.ClientRspHeader{}
	opts := []client.Option{
		client.WithProtocol(entity.ProtocolHTTP),
		client.WithReqHead(reqHead),
		client.WithRspHead(rsphead),
		client.WithServiceName(routerConf.Service),
		client.WithCalleeMethod(method),
		client.WithTarget(routerConf.Target),
		client.WithTimeout(timeout),
		//client.WithSerializationType(codec.SerializationTypeNoop),
		client.WithCurrentSerializationType(codec.SerializationTypeNoop), // 只透传不序列化
		client.WithCurrentCompressType(codec.CompressTypeNoop),
	}
	if err = client.DefaultClient.Invoke(ctx, reqBody, rspBody, opts...); err != nil {
		return err
	}
	if err = transHTTPRsp(w, rsphead, rspBody); err != nil {
		log.Errorf("transHTTPRsp failed, err:%v", err)
		return errs.New(entity.ErrTransHTTPRsp, "trans to http.rsp err")
	}
	return nil
}

func transHTTPReq(r *http.Request) (*thttp.ClientReqHeader, *codec.Body, error) {
	reqHead := &thttp.ClientReqHeader{
		Method: r.Method,
		Host:   r.Host,
		Header: r.Header,
	}
	reqBody := &codec.Body{}
	if r.Method == http.MethodPost {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Errorf("transHTTPReq ioutil.ReadAll failed, err:%v", err)
			return nil, nil, errs.New(entity.ErrReadHTTPReqBody, "read HTTP req body err")
		}
		reqBody.Data = data
	}
	return reqHead, reqBody, nil
}

func transHTTPRsp(w http.ResponseWriter, rspHead *thttp.ClientRspHeader, rspBody *codec.Body) error {
	if rspHead != nil && rspBody != nil {
		copyRspHead(w, rspHead.Response)
		if rspHead.Response.StatusCode != http.StatusOK {
			w.WriteHeader(rspHead.Response.StatusCode)
		}
	}
	if rspBody != nil {
		_, err := w.Write(rspBody.Data)
		if err != nil {
			log.Errorf("write response body err: %v", err)
			return err
		}
	}
	return nil
}

func copyRspHead(w http.ResponseWriter, rsp *http.Response) {
	for k, v := range rsp.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}
}
