package logic

import (
	"context"
	"github.com/Yamon955/ShortVideo/interaction/repo/mysql"
	"github.com/Yamon955/ShortVideo/protocol/interaction/pb"
	"github.com/redis/go-redis/v9"
)

type Handler interface {
	// HandleLike 点赞
	HandleLike(ctx context.Context, req *pb.LikeReq) (*pb.LikeRsp, error)
	// HandleFollow 关注
	HandleFollow(ctx context.Context, req *pb.FollowReq) (*pb.FollowRsp, error)
	// HandleCommit 评论
	HandleCommit(ctx context.Context, req *pb.CommitReq) (*pb.CommitRsp, error)
	// HandleReport 举报
	HandleReport(ctx context.Context, req *pb.ReportReq) (*pb.ReportRsp, error)
}

type handlerImpl struct {
	db  mysql.DBClient
	rdb *redis.Client
}
