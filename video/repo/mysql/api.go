// Package mysql 数据库操作包
package mysql

import (
	"context"

	"github.com/Yamon955/ShortVideo/comm/base"
)

// DBClient 数据库操作接口对象
type DBClient interface {
	// InsertVideo 插入视频
	InsertVideo(ctx context.Context, video *base.Video) error
}

// NewDBClient 新建一个数据库操作接口对象
func NewDBClient() DBClient {
	return &dbClient{
		db,
	}
}
