package main

import (
	"net/http"

	"github.com/Yamon955/ShortVideo/gateway/entity"
	"github.com/Yamon955/ShortVideo/gateway/repo/router"
	_ "github.com/Yamon955/ShortVideo/gateway/repo/router"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/filter"
	thttp "trpc.group/trpc-go/trpc-go/http"
	"trpc.group/trpc-go/trpc-go/log"
)

func init() {
	// 服务端路由拦截器
	filter.Register("router_filter", routerFilter(), nil)
	// 服务端鉴权拦截器
	filter.Register("auth_filter", authFilter(), nil)
}

func main() {
	s := trpc.NewServer()
	if err := router.Init(); err != nil {
		log.Fatal(err)
	}
	thttp.HandleFunc("*", handle)
	thttp.RegisterNoProtocolService(s)
	if err := s.Serve(); err != nil {
		panic(err)
	}
}

func handle(w http.ResponseWriter, r *http.Request) error {
	routeConf := entity.GetRouteConf(r.Context())
	if routeConf.Protocol == entity.ProtocolTRPC {
		if err := forwardTRPC(w, r, routeConf); err != nil {
			log.Errorf("farward trpc failed, err:%v", err)
			return err
		}
		return nil
	}
	if err := forwardHTTP(w, r, routeConf); err != nil {
		log.Errorf("farward http failed, err:%v", err)
		return err
	}
	return nil
}

// curl -X POST -d '{"a":"1.0","b":"1.0"}' -H "Content-Type:application/json" "http://127.0.0.1:10001/test/caculator" -v

//curl -X POST -d '{"msg":"hello"}' -H "Content-Type:application/json" "http://127.0.0.1:10001/test" -v
