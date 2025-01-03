package mysql

import (
	"context"

	"github.com/Yamon955/ShortVideo/base"
	"gorm.io/gorm"
	"trpc.group/trpc-go/trpc-go/log"
)

type dbClient struct {
	client *gorm.DB
}

// FindUserByUsername 根据用户名查询用户
func (d *dbClient) FindUserByUsername(ctx context.Context, username string) (*base.User, error) {
	// var user *model.User 未初始化，为 nil
	user := new(base.User) // or user := &model.User{}
	// users 表建立了 username 的唯一索引，用于根据用户名查找用户
	res := d.client.Where("username = ?", username).Take(user)
	if err := res.Error; err != nil {
		log.ErrorContextf(ctx, "FindUserByUsername failed, username:%s, err:%v", username, err)
		return user, err
	}
	return user, nil
}

// FindUserByUID 根据用户 ID 查询用户
func (d *dbClient) FindUserByUID(ctx context.Context, uid uint64) (*base.User, error) {
	user := new(base.User) // or user := &model.User{}
	res := d.client.Where("id = ?", uid).Take(user)
	if err := res.Error; err != nil {
		log.ErrorContextf(ctx, "FindUserByUID failed, uid:%d, err:%v", uid, err)
		return user, err
	}
	return user, nil
}

// CreateUser 插入新用户
func (d *dbClient) CreateUser(ctx context.Context, user *base.User) error {
	res := d.client.Model(&base.User{}).Create(&user)
	if err := res.Error; err != nil {
		log.ErrorContextf(ctx, "CreateUser failed, err:%v", err)
		return err
	}
	return nil
}

// UpdateUserInfo 更新 user 信息
func (d *dbClient) UpdateUserInfo(ctx context.Context, uid uint64, updateInfo map[string]interface{}) error {
	res := d.client.Model(&base.User{}).Where("id = ?", uid).Updates(updateInfo)
	if err := res.Error; err != nil {
		log.ErrorContextf(ctx, "UpdateUserInfo failed, uid:%d, err:%v", uid, err)
		return err
	}
	return nil
}

// GetPublishListCount 获取作品列表总数
func (d *dbClient) GetPublishListCount(ctx context.Context, uid uint64) (int64, error) {
	var count int64
	// videos 表建立了 uid 的索引，用于查找用户的作品列表
	res := d.client.Model(&base.Video{}).Where("uid = ?", uid).Count(&count)
	if err := res.Error; err != nil && err != gorm.ErrRecordNotFound {
		log.ErrorContextf(ctx, "GetPublishListCount failed, uid:%d, err:%v", uid, err)
	}
	return count, nil
}

// GetLikedListCount 获取作品列表总数
func (d *dbClient) GetLikedListCount(ctx context.Context, uid uint64) (int64, error) {
	var count int64
	// likes 表建立了 uid 的索引，用于查找用户的作品列表
	res := d.client.Model(&base.Like{}).Where("uid = ?", uid).Count(&count)
	if err := res.Error; err != nil && err != gorm.ErrRecordNotFound {
		log.ErrorContextf(ctx, "GetPublishListCount failed, uid:%d, err:%v", uid, err)
	}
	return count, nil
}

// GetCollectListCount 获取作品列表总数
func (d *dbClient) GetCollectListCount(ctx context.Context, uid uint64) (int64, error) {
	var count int64
	// collects 表建立了 uid 的索引，用于查找用户的作品列表
	res := d.client.Model(&base.Collect{}).Where("uid = ?", uid).Count(&count)
	if err := res.Error; err != nil && err != gorm.ErrRecordNotFound {
		log.ErrorContextf(ctx, "GetPublishListCount failed, uid:%d, err:%v", uid, err)
	}
	return count, nil
}
