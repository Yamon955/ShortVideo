package main

import (
	"context"

	"github.com/Yamon955/ShortVideo/comm/utils"
	"github.com/Yamon955/ShortVideo/protocol/video/pb"
	"github.com/Yamon955/ShortVideo/video/entity/conf"
	"github.com/Yamon955/ShortVideo/video/entity/def"
	"github.com/Yamon955/ShortVideo/video/repo/minio"
	"github.com/Yamon955/ShortVideo/video/repo/mysql"
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
	if err := minio.Init(); err != nil {
		panic(err)
	}
	if err := utils.SnowflakeInit(getMachineID()); err != nil {
		panic(err)
	}
	if err := mysql.Init(); err != nil {
		panic(err)
	}
}

func main() {
	s := trpc.NewServer()

	// 注册两个不同协议的服务
	pb.RegisterVideoService(s.Service("trpc.shortvideo.video.Video"), newVideoSvr())
	pb.RegisterPublishService(s.Service("trpc.shortvideo.video.Publish"), newVideoSvr())

	if err := s.Serve(); err != nil {
		panic(err)
	}
}

// getMachineID 获取机器 ID 用于雪花算法生成 uuid，使用 redis 的 incy 函数保证机器 id 不同
func getMachineID() uint16 {
	machineID := redis.GetClient().Incr(context.Background(), def.MachineID).Val()
	machineID %= 1024
	return uint16(machineID)
}
