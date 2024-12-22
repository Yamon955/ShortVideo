package mysql

import (
	"context"

	"github.com/Yamon955/ShortVideo/user/entity/model"
)

type DBClient interface {
	// FindUserByUsername 数据库根据用户名查询用户
	FindUserByUsername(ctx context.Context, username string) (*model.User, error)
	// CreateUser 数据库插入新用户
	CreateUser(ctx context.Context, user *model.User) error
}

// NewDBClient 新建一个数据库连接
func NewDBClient() DBClient {
	return &dbClient{

	}
}
