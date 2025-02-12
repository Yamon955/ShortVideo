// Package mysql 数据库操作包
package mysql

import (
	"context"

	"github.com/Yamon955/ShortVideo/comm/base"
)

// DBClient 数据库操作接口对象
type DBClient interface {
	// InsertVideo 插入视频信息
	InsertVideo(ctx context.Context, video *base.Video) error
	// SelectUserPublishVideos 查询用户发布视频信息，每次返回 10 条数据
	SelectUserPublishVideos(ctx context.Context, id uint64, uid uint64) ([]*base.Video, error)
	// SelectVideosByVids 通过 vids 批量查询对应视频信息
	SelectVideosByVids(ctx context.Context, vids []uint64) ([]*base.Video, error)
	// SelectVideoByVid 通过 vid 查询对应视频信息
	SelectVideoByVid(ctx context.Context, vids []uint64) (*base.Video, error)
}

// NewDBClient 新建一个数据库操作接口对象
func NewDBClient() DBClient {
	return &dbClient{
		db,
	}
}
