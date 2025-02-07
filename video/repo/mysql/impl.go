package mysql

import (
	"context"

	"github.com/Yamon955/ShortVideo/comm/base"
	"gorm.io/gorm"
	"trpc.group/trpc-go/trpc-go/log"
)

type dbClient struct {
	client *gorm.DB
}

func (d *dbClient) InsertVideo(ctx context.Context, video *base.Video) error {
	res := d.client.Model(&base.Video{}).Create(video)
	if err := res.Error; err != nil {
		log.ErrorContextf(ctx, "InsertVideo failed, err:%v", err)
		return err
	}
	return nil
}
