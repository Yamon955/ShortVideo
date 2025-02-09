package logic

import (
	"context"
	"strconv"

	"github.com/Yamon955/ShortVideo/protocol/video/pb"
	"github.com/Yamon955/ShortVideo/video/entity/def"
	"github.com/Yamon955/ShortVideo/video/entity/errcode"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/errs"
)

// HandleGetPublishList 分页获取用户发布列表
func (h *handlerImpl) HandleGetPublishList(ctx context.Context, req *pb.GetPublishListReq) (*pb.GetPublishListRsp, error) {
	if req.GetStartId() < 0 {
		return nil, errs.New(errcode.ErrInvalidStartID, "invalid start_id")
	}
	uid, _ := strconv.ParseUint(string(trpc.GetMetaData(ctx, "sv_login_uid")), 10, 64)
	videos, err := h.DB.SelectUserPublishVideos(ctx, req.GetStartId(), uid)
	if err != nil {
		return nil, errs.New(errcode.ErrDBOperation, "db operation failed")
	}
	var nextId uint64
	isFinash := true
	if len(videos) >= def.SQLLimitSize {
		nextId = videos[def.SQLLimitSize-1].ID
		isFinash = false
	}
	videoInfos := make([]*pb.VideoInfo, 0, def.SQLLimitSize)
	for _, video := range videos {
		videoInfo := &pb.VideoInfo{
			Vid:           video.VID,
			VideoUrl:      video.VideoURL,
			CoverUrl:      video.CoverURL,
			Title:         video.Title,
			Tags:          video.Tags,
			LikesCount:    video.LikesCount,
			CollectCount:  video.CollectsCount,
			CommentsCount: video.CommentsCount,
			PublishTime:   video.PublishTime,
		}
		videoInfos = append(videoInfos, videoInfo)
	}
	return &pb.GetPublishListRsp{
		NextId:     nextId,
		IsFinash:   isFinash,
		VideoInfos: videoInfos,
	}, nil
}
