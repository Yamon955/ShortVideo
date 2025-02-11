package main

import (
	"github.com/Yamon955/ShortVideo/protocol/recommend/pb"
	"github.com/Yamon955/ShortVideo/recommend/entity/conf"
	"github.com/Yamon955/ShortVideo/recommend/repo/redis"
	"trpc.group/trpc-go/trpc-go"
)

func main() {
	s := trpc.NewServer()

	if err := conf.Init(); err != nil {
		panic(err)
	}
	if err := redis.Init(); err != nil {
		panic(err)
	}

	pb.RegisterRecommendService(s, newRecommendSvr())

	if err := s.Serve(); err != nil {
		panic(err)
	}
}
