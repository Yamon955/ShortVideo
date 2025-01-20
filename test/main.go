package main

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"mime/multipart"
	"time"

	"github.com/Yamon955/ShortVideo/test/pb"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/codec"
	thttp "trpc.group/trpc-go/trpc-go/http"
	"trpc.group/trpc-go/trpc-go/log"
)

const pieceSize = 1024 * 1024

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
	rsp := &pb.HelloResponse{
		Msg: "success",
	}
	return rsp, nil
}

// parseHttpForm 解析http_form请求
func parseHttpForm(ctx context.Context) (err error) {
	head := thttp.Head(ctx)
	head.Request.Body = ioutil.NopCloser(bytes.NewReader(head.ReqBody))
	err = head.Request.ParseMultipartForm(1) // 32MB
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
		//log.InfoContextf(ctx, "fileHeader:%+v", fileHeader)
		log.ErrorContextf(ctx, "file:%v", file)
		//_, _ = pieceDownloadFeed(ctx, file, fileHeader)
		_, _ = downloadFeed(ctx, file, fileHeader)
	}
	return
}

// pieceDownloadFeed 分片读取视频数据
func pieceDownloadFeed(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) ([]byte, error) {
	st := time.Now()
	pieceNum := fileHeader.Size / pieceSize
	if fileHeader.Size%pieceSize != 0 {
		pieceNum += 1
	}
	data := make([][]byte, pieceNum)
	var handlers []func() error
	for i := 0; i < int(pieceNum); i++ {
		index := i
		start := int64(index * pieceSize)
		partSize := int64(pieceSize)
		if i == int(pieceNum)-1 {
			partSize = pieceSize - 1
		}
		handlers = append(handlers, func() error {
			sectionReader := io.NewSectionReader(file, start, partSize)
			buffer := make([]byte, partSize)
			_, err := sectionReader.Read(buffer)
			if err != nil && err != io.EOF {
				return err
			}
			data[index] = buffer
			return nil
		})
	}
	if err := trpc.GoAndWait(handlers...); err != nil {
		log.ErrorContextf(ctx, "pieceDownloadFeed failed, err:%v", err)
		return nil, err
	}
	res := make([]byte, 0, fileHeader.Size)
	for _, partData := range data {
		res = append(res, partData...)
	}
	elapsed := time.Since(st)
	log.Infof("函数执行时间: %d毫秒", elapsed.Milliseconds())
	size := float64(len(res) * 1.0 / 1024 / 1024)
	log.Infof("len(res) = %fMB", size)
	log.Infof("res:%v", res[:pieceSize/256])
	return res, nil
}

func downloadFeed(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) ([]byte, error) {
	st := time.Now()
	// 使用缓冲区逐块读取文件内容并写入 req.Data
	var buffer bytes.Buffer
	_, err := io.Copy(&buffer, file)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}
	res := buffer.Bytes()
	elapsed := time.Since(st)
	log.Infof("函数执行时间: %d毫秒", elapsed.Milliseconds())
	size := float64(len(res) * 1.0 / 1024 / 1024)
	log.Infof("len(res) = %fMB", size)
	log.Infof("res:%v", res[:pieceSize/256])
	return res, nil
}
