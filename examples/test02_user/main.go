package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Yamon955/ShortVideo/user/pb"
	"trpc.group/trpc-go/trpc-go/client"
)

func main() {
	opts := []client.Option{
		client.WithTarget("ip://localhost:20001"),
		client.WithTimeout(time.Second * 2),
	}
	proxy := pb.NewUserClientProxy(opts...)
	test_login(proxy)
	//test_register(proxy)
	//test_login(proxy)

}

func test_login(proxy pb.UserClientProxy) {
	req := &pb.LoginReq{
		Username: "adin",
		Password: "12346788",
	}
	rsp, err := proxy.Login(context.Background(), req)
	if err != nil {
		fmt.Printf("login failed, err:%v\n", err)
		return
	}
	fmt.Printf("login suc, rsp:%v\n", rsp)
}

func test_register(proxy pb.UserClientProxy) {
	req := &pb.RegisterReq{
		Username: "admin",
		Password: "12345678",
	}
	rsp, err := proxy.Register(context.Background(), req)
	if err != nil {
		fmt.Printf("register failed, err:%v\n", err)
		return
	}
	fmt.Printf("register suc, rsp:%v\n", rsp)
}
