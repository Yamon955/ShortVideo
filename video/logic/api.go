package logic

import (
	"context"

	"github.com/Yamon955/ShortVideo/protocol/video/pb"
	"github.com/Yamon955/ShortVideo/video/repo/mysql"
	"github.com/Yamon955/ShortVideo/video/repo/uploader"
)

// Handler 视频服务处理器
type Handler interface {
	// HandlePublish 视频发布
	HandlePublish(ctx context.Context, req *pb.PublishReq) (*pb.PublishRsp, error)
	// HandleGetFeeds 获取推荐视频列表
	HandleGetFeeds(ctx context.Context, req *pb.GetFeedsReq) (*pb.GetFeedsRsp, error)
	// HandleGetPublishList 分页获取用户发布列表
	HandleGetPublishList(ctx context.Context, req *pb.GetPublishListReq) (*pb.GetPublishListRsp, error)
}

func NewHandler() Handler {
	return &handlerImpl{
		Uploader: uploader.NewMediaFileUploader(),
		DB:       mysql.NewDBClient(),
	}
}

type handlerImpl struct {
	Uploader uploader.MediaFileUploader
	DB       mysql.DBClient
}
