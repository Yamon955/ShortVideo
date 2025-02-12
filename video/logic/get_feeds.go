package logic

import (
	"context"
	"strconv"

	recommendpb "github.com/Yamon955/ShortVideo/protocol/recommend/pb"
	"github.com/Yamon955/ShortVideo/protocol/video/pb"
	"github.com/Yamon955/ShortVideo/video/entity/errcode"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/errs"
	"trpc.group/trpc-go/trpc-go/log"
)

// HandleGetFeeds 获取推荐视频列表
func (h *handlerImpl) HandleGetFeeds(ctx context.Context, req *pb.GetFeedsReq) (*pb.GetFeedsRsp, error) {
	uid, _ := strconv.ParseUint(string(trpc.GetMetaData(ctx, "sv_login_uid")), 10, 64)
	proxy := recommendpb.NewRecommendClientProxy()
	recommendReq := &recommendpb.FeedsRecommendReq{
		Uid:        uid,
		StartIndex: 0,
	}
	recommendRsp, err := proxy.FeedsRecommend(ctx, recommendReq)
	if err != nil {
		log.ErrorContextf(ctx, "FeedsRecommend failed, err:%v", err)
		return nil, err
	}
	// 获取视频信息
	log.Infof("recommendRsp %+v", recommendRsp.GetVids())
	vidInfos, err := h.DB.SelectVideosByVids(ctx, recommendRsp.GetVids())
	log.Infof("vidInfos %+v", vidInfos)
	if err != nil {
		log.ErrorContextf(ctx, "SelectVideosByVids failed, err:%v", err)
		return nil, errs.New(errcode.ErrDBOperation, "db operation failed")
	}
	videoInfos := make([]*pb.VideoInfo, 0, len(vidInfos))
	for _, vidInfo := range vidInfos {
		videoInfos = append(videoInfos, &pb.VideoInfo{
			Vid:           vidInfo.VID,
			CoverUrl:      vidInfo.CoverURL,
			Title:         vidInfo.Title,
			Tags:          vidInfo.Tags,
			LikesCount:    vidInfo.LikesCount,
			CollectCount:  vidInfo.CollectsCount,
			CommentsCount: vidInfo.CommentsCount,
			PublishTime:   vidInfo.PublishTime,
		})
	}
	return &pb.GetFeedsRsp{
		NextIndex:  recommendRsp.GetNextIndex(),
		VideoInfos: videoInfos,
	}, nil
}
