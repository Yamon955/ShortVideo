package entity

import "context"

// 使用内置类型（如 string）作为上下文键可能会导致键冲突，因此定义一个自定义类型用于上下文键
type contextKey string

const RouterKey contextKey = "SHORTVIDEO_GATEWAY_ROUTER_KEY"

// WithRouteConf is used to set route config in ctx
func WithRouteConf(ctx context.Context, routeConf *RouteConf) context.Context {
	return context.WithValue(ctx, RouterKey, routeConf)
}

// GetRouteConf is used to get router cofig from ctx
func GetRouteConf(ctx context.Context) *RouteConf {
	if router, ok := ctx.Value(RouterKey).(*RouteConf); ok {
		return router
	}
	return nil
}
