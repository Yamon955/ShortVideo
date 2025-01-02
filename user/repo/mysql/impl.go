package mysql

import (
	"context"

	"github.com/Yamon955/ShortVideo/user/entity/model"
	"gorm.io/gorm"
)

type dbClient struct {
	client *gorm.DB
}

// FindUserByUsername 根据用户名查询用户
func (d *dbClient) FindUserByUsername(_ context.Context, username string) (*model.User, error) {
	// var user *model.User 未初始化，为 nil
	user := new(model.User) // or user := &model.User{}
	res := d.client.Where("username = ?", username).Take(user)
	return user, res.Error
}

// FindUserByUID 根据用户 ID 查询用户
func (d *dbClient) FindUserByUID(_ context.Context, uid uint64) (*model.User, error) {
	user := new(model.User) // or user := &model.User{}
	res := d.client.Where("id = ?", uid).Take(user)
	return user, res.Error
}

// CreateUser 插入新用户
func (d *dbClient) CreateUser(_ context.Context, user *model.User) error {
	res := d.client.Model(&model.User{}).Create(&user)
	return res.Error
}
