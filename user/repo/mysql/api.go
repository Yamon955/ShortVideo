// Package mysql 数据库操作包
package mysql

import (
	"context"

	"github.com/Yamon955/ShortVideo/base"
)

// DBClient 数据库操作接口对象
type DBClient interface {
	// FindUserByUsername 根据用户名查询用户
	FindUserByUsername(ctx context.Context, username string) (*base.User, error)
	// FindUserByUID 根据用户 ID 查询用户
	FindUserByUID(_ context.Context, uid uint64) (*base.User, error)
	// CreateUser 插入新用户
	CreateUser(ctx context.Context, user *base.User) error
	// UpdateUserInfo 更新 user 信息
	UpdateUserInfo(_ context.Context, uid uint64, updateInfo map[string]interface{}) error
	// GetPublishListCount 获取作品列表总数
	GetPublishListCount(ctx context.Context, uid uint64) (int64, error)
	// GetLikedListCount 获取作品列表总数
	GetLikedListCount(ctx context.Context, uid uint64) (int64, error)
	// GetCollectListCount 获取作品列表总数
	GetCollectListCount(ctx context.Context, uid uint64) (int64, error)
}

// NewDBClient 新建一个数据库操作接口对象
func NewDBClient() DBClient {
	return &dbClient{
		db,
	}
}
