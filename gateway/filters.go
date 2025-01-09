package main

import (
	"context"

	"github.com/Yamon955/ShortVideo/comm/utils"
	"github.com/Yamon955/ShortVideo/gateway/entity"
	"github.com/Yamon955/ShortVideo/gateway/repo/auth"
	"github.com/Yamon955/ShortVideo/gateway/repo/router"
	"trpc.group/trpc-go/trpc-go/errs"
	"trpc.group/trpc-go/trpc-go/filter"
	thttp "trpc.group/trpc-go/trpc-go/http"
)

// routerFilter 路由 filter 需要放在依赖 entity.Router 的 filter 前面
func routerFilter() filter.ServerFilter {
	return func(ctx context.Context, req interface{}, next filter.ServerHandleFunc) (rsp interface{}, err error) {
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

// authFilter 权限拦截器，校验登录态
func authFilter() filter.ServerFilter {
	return func(ctx context.Context, req interface{}, next filter.ServerHandleFunc) (rsp interface{}, err error) {
		routeConf := entity.GetRouteConf(ctx)
		isAllowNoAuth := false
		if routeConf != nil {
			isAllowNoAuth = routeConf.IsAllowNoAuth
		}
		// 允许无登录态请求的服务接口，无需校验，直接放行
		if isAllowNoAuth {
			return next(ctx, req)
		}
		token := thttp.Head(ctx).Request.Header.Get(utils.SignKey)
		verifySuc, err := auth.VerifyJWT(ctx, token)
		if err != nil || !verifySuc {
			return rsp, errs.New(entity.ErrVerifAuthFail, "invalid token")
		}
		return next(ctx, req)
	}
}
