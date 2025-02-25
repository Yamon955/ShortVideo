package logic

import (
	"context"

	"github.com/Yamon955/ShortVideo/protocol/video/pb"
	"github.com/Yamon955/ShortVideo/video/repo/mysql"
	MyRedis "github.com/Yamon955/ShortVideo/video/repo/redis"
	"github.com/Yamon955/ShortVideo/video/repo/uploader"
	"github.com/redis/go-redis/v9"
)

// Handler 视频服务处理器
type Handler interface {
	// HandlePublish 视频发布
	HandlePublish(ctx context.Context, req *pb.PublishReq) (*pb.PublishRsp, error)
	// HandleGetFeeds 获取推荐视频列表
	HandleGetFeeds(ctx context.Context, req *pb.GetFeedsReq) (*pb.GetFeedsRsp, error)
	// HandleGetPublishList 分页获取用户发布列表
	HandleGetPublishList(ctx context.Context, req *pb.GetPublishListReq) (*pb.GetPublishListRsp, error)
	// HandleWatchVideo 处理用户观看的 video，加入到布隆过滤器设置为已观看
	HandleWatchVideo(ctx context.Context, req *pb.WatchVideoReq) (*pb.WatchVideoRsp, error)
}

func NewHandler() Handler {
	return &handlerImpl{
		uploader: uploader.NewMediaFileUploader(),
		db:       mysql.NewDBClient(),
		rdb:      MyRedis.GetClient(),
	}
}

type handlerImpl struct {
	uploader uploader.MediaFileUploader
	db       mysql.DBClient
	rdb      *redis.Client
}
