// Package mysql 数据库操作包
package mysql

import (
	"context"

	"github.com/Yamon955/ShortVideo/user/entity/model"
)

// DBClient 数据库操作接口对象
type DBClient interface {
	// FindUserByUsername 根据用户名查询用户
	FindUserByUsername(ctx context.Context, username string) (*model.User, error)
	// CreateUser 插入新用户
	CreateUser(ctx context.Context, user *model.User) error
}

// NewDBClient 新建一个数据库操作接口对象
func NewDBClient() DBClient {
	return &dbClient{
		db,
	}
}
