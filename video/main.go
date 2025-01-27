package main

import (
	"github.com/Yamon955/ShortVideo/protocol/video/pb"
	"github.com/Yamon955/ShortVideo/video/entity/conf"
	"github.com/Yamon955/ShortVideo/video/repo/redis"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/codec"
)

func init() {
	// 对于 Content-Type=multipart/form-data 的请求无需反序列化，因此注册对应的反序列化处理对象，重写其 Unmarshal 方法
	codec.RegisterSerializer(SerializationTypeFormData, &FormDataSerialization{})

	if err := conf.Init(); err != nil {
		panic(err)
	}
	if err := redis.Init(); err != nil {
		panic(err)
	}
}

func main() {
	s := trpc.NewServer()
	pb.RegisterVideoService(s, &videoSvrImpl{})

	if err := s.Serve(); err != nil {
		panic(err)
	}
}
