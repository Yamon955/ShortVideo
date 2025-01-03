package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Yamon955/ShortVideo/protocol/user/pb"
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
	test_setProfile(proxy)
	test_getProfile(proxy)

}

func test_login(proxy pb.UserClientProxy) {
	req := &pb.LoginReq{
		Username: "admin",
		Password: "12345678",
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

func test_getProfile(proxy pb.UserClientProxy) {
	req := &pb.BatchGetProfileReq{
		Uids:         []uint64{1, 2},
		ProfileTypes: []pb.PROFILE_TYPES{pb.PROFILE_TYPES_MAIN_PAGE_INFO, pb.PROFILE_TYPES_PUBLISH_LIST_COUNT},
	}
	rsp, err := proxy.BatchGetProfile(context.Background(), req)
	if err != nil {
		fmt.Printf("get profile failed, err:%v\n", err)
		return
	}
	fmt.Printf("get profile suc, rsp:%v\n", rsp)
}

func test_setProfile(proxy pb.UserClientProxy) {
	req := &pb.SetProfileReq{
		Uid:          1,
		ProfileTypes: []pb.PROFILE_TYPES{pb.PROFILE_TYPES_USERNAME},
		Username:     "admin",
	}
	rsp, err := proxy.SetProfile(context.Background(), req)
	if err != nil {
		fmt.Printf("set profile failed, err:%v\n", err)
		return
	}
	fmt.Printf("set profile suc, rsp:%v\n", rsp)
}
