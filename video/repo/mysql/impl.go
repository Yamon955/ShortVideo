package mysql

import (
	"context"

	"github.com/Yamon955/ShortVideo/comm/base"
	"github.com/Yamon955/ShortVideo/video/entity/def"
	"gorm.io/gorm"
	"trpc.group/trpc-go/trpc-go/log"
)

type dbClient struct {
	client *gorm.DB
}

// InsertVideo 插入视频信息
func (d *dbClient) InsertVideo(ctx context.Context, video *base.Video) error {
	res := d.client.Model(&base.Video{}).Create(video)
	if err := res.Error; err != nil {
		log.ErrorContextf(ctx, "InsertVideo failed, err:%v", err)
		return err
	}
	return nil
}

// SelectUserPublishVideos 查询用户发布视频信息，每次返回 10 条数据
func (d *dbClient) SelectUserPublishVideos(ctx context.Context, id uint64, uid uint64) ([]*base.Video, error) {
	videos := make([]*base.Video, def.SQLLimitSize)
	res := d.client.Where("id > ? and uid = ?", id, uid).Limit(def.SQLLimitSize).Find(&videos)
	if err := res.Error; err != nil {
		log.ErrorContextf(ctx, "SelectUserPublishVideos failed, err:%v", err)
		return nil, err
	}
	return videos, nil
}

// SelectVideosByVids 通过 vids 批量查询对应视频信息
func (d *dbClient) SelectVideosByVids(ctx context.Context, vids []uint64) ([]*base.Video, error) {
	videoInfos := make([]*base.Video, 0, len(vids))
	res := d.client.Where("vid IN ?", vids).Find(&videoInfos)
	if err := res.Error; err != nil {
		log.ErrorContextf(ctx, "SelectVideosByVids failed, err:%v", err)
		return nil, err
	}
	return videoInfos, nil
}

// SelectVideoByVid 通过 vid 查询对应视频信息
func (d *dbClient) SelectVideoByVid(ctx context.Context, vids []uint64) (*base.Video, error) {
	videoInfo := new(base.Video)
	res := d.client.Where("vid IN ?", vids).Find(&videoInfo)
	if err := res.Error; err != nil {
		log.ErrorContextf(ctx, "SelectVideosByVids failed, err:%v", err)
		return nil, err
	}
	return videoInfo, nil
}
