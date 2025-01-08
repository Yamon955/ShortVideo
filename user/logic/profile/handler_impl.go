package profile

import (
	"context"

	"github.com/Yamon955/ShortVideo/base"
	"github.com/Yamon955/ShortVideo/protocol/user/pb"
	"github.com/Yamon955/ShortVideo/user/entity/def"
	"github.com/Yamon955/ShortVideo/user/repo/mysql"
	"gorm.io/gorm"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
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
	// 默认拉取个人主页资料
	if len(req.GetProfileTypes()) == 0 {
		req.ProfileTypes = append(req.GetProfileTypes(), pb.PROFILE_TYPES_MAIN_PAGE_INFO)
	}
	for _, uid := range req.GetUids() {
		if uid < def.MIN_UID {
			rsp.FailedUIDs = append(rsp.GetFailedUIDs(), uid)
			log.ErrorContextf(ctx, "uid:%d is invalid", uid)
			continue
		}
		var (
			handlers         []func() error
			user             = &base.User{}
			publishListCount int64
			likedListCount   int64
			collectListCount int64
		)
		for _, infoType := range req.GetProfileTypes() {
			switch infoType {
			case pb.PROFILE_TYPES_MAIN_PAGE_INFO:
				handlers = append(handlers, func() (err error) {
					user, err = h.db.FindUserByUID(ctx, uid)
					return
				})
			case pb.PROFILE_TYPES_PUBLISH_LIST_COUNT:
				handlers = append(handlers, func() (err error) {
					publishListCount, err = h.db.GetPublishListCount(ctx, uid)
					return err
				})
			case pb.PROFILE_TYPES_LIKED_LIST_COUNT:
				handlers = append(handlers, func() (err error) {
					likedListCount, err = h.db.GetLikedListCount(ctx, uid)
					return err
				})
			case pb.PROFILE_TYPES_COLLECT_LIST_COUNT:
				handlers = append(handlers, func() (err error) {
					collectListCount, err = h.db.GetCollectListCount(ctx, uid)
					return err
				})
			}
		}
		// 并行拉取
		if err := trpc.GoAndWait(handlers...); err != nil {
			log.ErrorContextf(ctx, "GoAndWait failed, uid:%d, err:%v", uid, err)
			rsp.FailedUIDs = append(rsp.GetFailedUIDs(), uid)
			continue
		}
		log.Infof("user:%v, publisListCount:%d, likedListCount:%d, collectListCount:%d", user, publishListCount, likedListCount, collectListCount)
		userInfo := fillUserInfo(user, publishListCount, likedListCount, collectListCount)
		rsp.UserInfos = append(rsp.UserInfos, userInfo)
	}
	return nil
}

// HandleSetProfile 处理用户资料修改
func (h *handlerImpl) HandleSetProfile(
	ctx context.Context,
	req *pb.SetProfileReq,
	rsp *pb.SetProfileRsp) error {
	// 要更新的信息
	updateInfo := make(map[string]interface{})
	// 更新失败的字段
	failedTypes := make(map[int32]string)
	for _, infoType := range req.GetProfileTypes() {
		switch infoType {
		case pb.PROFILE_TYPES_USERNAME:
			if checkUsername(ctx, req.GetUsername(), failedTypes, h.db) {
				updateInfo["username"] = req.GetUsername()
			}
		case pb.PROFILE_TYPES_AVATOR:
			updateInfo["avator"] = req.GetAvator()
		case pb.PROFILE_TYPES_SIGN:
			updateInfo["sign"] = req.GetSign()
		case pb.PROFILE_TYPES_GENDER:
			updateInfo["gender"] = req.GetGender()
		}
	}
	// 更新用户信息
	if err := h.db.UpdateUserInfo(ctx, req.GetUid(), updateInfo); err != nil {
		for _, v := range req.GetProfileTypes() {
			_, exist := failedTypes[int32(v)]
			if !exist {
				failedTypes[int32(v)] = err.Error()
			}
		}
	}
	rsp.FailedTypes = failedTypes
	return nil
}

func fillUserInfo(user *base.User, publishListCount int64, likedListCount int64, collectListCount int64) *pb.UserInfo {
	userInfo := &pb.UserInfo{}
	if user != nil {
		userInfo = &pb.UserInfo{
			MainPageInfo: &pb.MainPageInfo{
				ID:       user.ID,
				Username: user.Username,
				Avator:   user.Avatar,
				Sign:     user.Sign,
				Gender:   user.Gender,
			},
		}
	}
	userInfo.PublishListCount = int32(publishListCount)
	userInfo.LikedListCount = int32(likedListCount)
	userInfo.CollectListCount = int32(collectListCount)
	return userInfo
}

// checkUsername 检查用户名是否合法
func checkUsername(ctx context.Context, username string, failedTypes map[int32]string, db mysql.DBClient) bool {
	if len(username) == 0 {
		failedTypes[int32(pb.PROFILE_TYPES_USERNAME)] = "username isn't null"
		return false
	}
	if len(username) > def.MAX_LEN {
		failedTypes[int32(pb.PROFILE_TYPES_USERNAME)] = "username too long"
		return false
	}
	_, err := db.FindUserByUsername(ctx, username)
	if err == nil {
		failedTypes[int32(pb.PROFILE_TYPES_USERNAME)] = "username is used"
		return false
	} else if err != gorm.ErrRecordNotFound {
		failedTypes[int32(pb.PROFILE_TYPES_USERNAME)] = err.Error()
		return false
	}
	return true
}
