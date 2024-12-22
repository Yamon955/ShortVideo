package router

import (
	"context"

	"github.com/Yamon955/ShortVideo/gateway/entity"
)

// Reader 路由配置读取接口
type Reader interface {
	// GetMatchRouter 匹配路由
	GetMatchRouter(ctx context.Context, path string) (*entity.RouteConf, error)
}

// NewReader 新建路由配置读取接口
func NewReader() Reader {
	return newReaderImpl()
}