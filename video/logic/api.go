package logic

import (
	"context"

	"github.com/Yamon955/ShortVideo/protocol/video/pb"
	"github.com/Yamon955/ShortVideo/video/repo/mysql"
	"github.com/Yamon955/ShortVideo/video/repo/uploader"
)

type Handler interface {
	HandleGetFeeds(ctx context.Context, req *pb.GetFeedsReq) (*pb.GetFeedsRsp, error)
	HandlePublish(ctx context.Context, req *pb.PublishReq) (*pb.PublishRsp, error)
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
