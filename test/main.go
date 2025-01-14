package main

import (
	"bytes"
	"context"
	"io/ioutil"

	"github.com/Yamon955/ShortVideo/test/pb"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/codec"
	thttp "trpc.group/trpc-go/trpc-go/http"
	"trpc.group/trpc-go/trpc-go/log"
)

func init() {
	SerializationTypeFormData := 131
	// 对于 Content-Type=multipart/form-data 的请求无需反序列化，因此注册对应的反序列化处理对象，重写其 Unmarshal 方法
	codec.RegisterSerializer(SerializationTypeFormData, &FormDataSerialization{})
}

type FormDataSerialization struct {
	codec.JSONPBSerialization
}

// Unmarshal multipart/form-data 类型解码方法, 重写 Unmarshal 方法，对于 form-data 数据无需反序列化
func (s *FormDataSerialization) Unmarshal(in []byte, body interface{}) error {
	return nil
}

func main() {
	s := trpc.NewServer()
	pb.RegisterHelloWorldServiceService(s, &helloWorldServiceImpl{})
	if err := s.Serve(); err != nil {
		panic(err)
	}
}

type helloWorldServiceImpl struct{}

func (s *helloWorldServiceImpl) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	_ = parseHttpForm(ctx)
	data := req.Data
	log.Infof("len:%d", len(data))
	log.Infof("req:%v", req)
	rsp := &pb.HelloResponse{
		Msg: "success",
	}
	return rsp, nil
}

// parseHttpForm 解析http_form请求
func parseHttpForm(ctx context.Context) (err error) {
	head := thttp.Head(ctx)
	head.Request.Body = ioutil.NopCloser(bytes.NewReader(head.ReqBody))

	err = head.Request.ParseMultipartForm(32 << 20) // 32M
	if err != nil {
		log.ErrorContextf(ctx, "ParseMultipartForm failed, err:%v", err)
		return err
	}

	// 获取 form-data 值参数（kv结构）
	//formValues := head.Request.MultipartForm.Value

	// 获取 form-data 文件参数
	formFiles := head.Request.MultipartForm.File
	for formName := range formFiles {
		file, fileHeader, err := head.Request.FormFile(formName) // file = fileHeader.Open()
		if nil != err {
			log.ErrorContextf(ctx, "read FormFile failed, err:%v:", err)
			continue
		}
		log.InfoContextf(ctx, "fileHeader:%+v", fileHeader)
		log.ErrorContextf(ctx, "file:%v", file)
	}
	return
}
