package main

import (
	"context"

	"github.com/Yamon955/ShortVideo/gateway/entity"
	"github.com/Yamon955/ShortVideo/gateway/repo/router"
	"trpc.group/trpc-go/trpc-go/filter"
	thttp "trpc.group/trpc-go/trpc-go/http"
)


// routerFilter 路由 filter 需要放在依赖 entity.Router 的 filter 前面
func routerFilter() filter.ServerFilter {
	return func(ctx context.Context, req interface{}, next filter.ServerHandleFunc) (rsp interface{}, err error){
		routerReader := router.NewReader()
		header := thttp.Head(ctx)
		reqPath := header.Request.URL.Path
		matchRouter, err := routerReader.GetMatchRouter(ctx, reqPath)
		if err != nil {
			return nil, err
		}
		return next(entity.WithRouteConf(ctx, matchRouter), req)
	}
}