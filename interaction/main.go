package interaction

import (
	"github.com/Yamon955/ShortVideo/interaction/entity/conf"
	"github.com/Yamon955/ShortVideo/interaction/repo/mysql"
	"github.com/Yamon955/ShortVideo/interaction/repo/redis"
	"github.com/Yamon955/ShortVideo/protocol/interaction/pb"
	"trpc.group/trpc-go/trpc-go"
)

func main() {
	s := trpc.NewServer()

	// 初始化配置
	if err := conf.Init(); err != nil {
		panic(err)
	}
	if err := mysql.Init(); err != nil {
		panic(err)
	}
	if err := redis.Init(); err != nil {
		panic(err)
	}

	pb.RegisterInteractionService(s, newInteractionSvr())

	if err := s.Serve(); err != nil {
		panic(err)
	}
}
