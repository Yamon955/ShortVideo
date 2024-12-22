package main

import (
	"github.com/Yamon955/ShortVideo/user/entity/conf"
	"github.com/Yamon955/ShortVideo/user/pb"
	"trpc.group/trpc-go/trpc-go"
)

func main() {
	s := trpc.NewServer()

	// 初始化配置
	if err := conf.Init(); err != nil {
		panic(err)
	}

	pb.RegisterUserService(s, newUserSvr())

	if err := s.Serve(); err != nil {
		panic(err)
	}
	
}