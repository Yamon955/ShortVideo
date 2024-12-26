package mysql

import (
	"context"

	"github.com/Yamon955/ShortVideo/user/entity/model"
	"gorm.io/gorm"
)

type dbClient struct {
	client *gorm.DB
}

func (d *dbClient) FindUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user *model.User
	err := d.client.Where("username = ?", username).Take(user).Error
	return user, err
}

func (d *dbClient) CreateUser(ctx context.Context, user *model.User) error {
	err := d.client.Model(&model.User{}).Create(user).Error
	return err
}
