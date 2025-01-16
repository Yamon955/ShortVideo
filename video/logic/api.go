package logic

import (
	"context"

	"github.com/Yamon955/ShortVideo/protocol/video/pb"
	"github.com/Yamon955/ShortVideo/video/repo/downloader"
	"github.com/Yamon955/ShortVideo/video/repo/uploader"
)

type Handler interface {
	HandleGetFeeds(ctx context.Context, req *pb.GetFeedsReq) (*pb.GetFeedsRsp, error)
	HandlePublish(ctx context.Context, req *pb.PublishReq) (*pb.PublishRsp, error)
	HandleGetPublishList(ctx context.Context, req *pb.GetPublishListReq) (*pb.GetPublishListRsp, error)
}

func NewHandler() Handler {
	return &handlerImpl{
		Downloader: downloader.NewMediaFileDownloader(),
	}
}

type handlerImpl struct {
	Downloader downloader.MediaFileDownloader
	Uploader   uploader.MediaFileUploader
}
