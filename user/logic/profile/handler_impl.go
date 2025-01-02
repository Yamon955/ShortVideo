package profile

import (
	"context"

	"github.com/Yamon955/ShortVideo/protocol/user/pb"
	"github.com/Yamon955/ShortVideo/user/entity/model"
	"github.com/Yamon955/ShortVideo/user/repo/mysql"
	"trpc.group/trpc-go/trpc-go/log"
)

const (
	MinUID = 0
)

// handlerImpl 用户资料处理器实现
type handlerImpl struct {
	db mysql.DBClient
}

// HandleBatchGetProfile 处理用户资料获取
func (h *handlerImpl) HandleBatchGetProfile(
	ctx context.Context,
	req *pb.BatchGetProfileReq,
	rsp *pb.BatchGetProfileRsp) error {
	for _, uid := range req.GetUid() {
		if uid <= MinUID {
			log.ErrorContextf(ctx, "uid:%d is invalid", uid)
			continue
		}
		user, err := h.db.FindUserByUID(ctx, uid)
		if err != nil {
			log.ErrorContextf(ctx, "uid:%d GetProfile failed, err:%v", uid, err)
			continue
		}
		userInfo := fillUserInfo(user)
		rsp.UserInfos = append(rsp.UserInfos, userInfo)
	}
	return nil
}

// HandleSetProfile 处理用户资料修改
func (h *handlerImpl) HandleSetProfile(
	ctx context.Context,
	req *pb.SetProfileReq,
	rsp *pb.SetProfileRsp) error {
	updateInfo := make(map[string]interface{})
	for _, infoType := range req.GetProfileTypes() {
		if infoType == pb.PROFILE_TYPES_all {
			updateInfo["username"]
		}
		switch infoType {
		case pb.PROFILE_TYPES_all:

		}
	}
	return nil
}

func fillUserInfo(user *model.User) (userInfo *pb.UserInfo) {
	if user == nil {
		return nil
	}
	userInfo = &pb.UserInfo{
		ID:           user.ID,
		Username:     user.Username,
		Avator:       user.Avatar,
		Sign:         user.Sign,
		Gender:       int32(user.Gender),
		FansCount:    int32(user.FansCount),
		FollowsCount: int32(user.FollowsCount),
		RegisterTime: user.RegisterTime,
	}
	return
}
