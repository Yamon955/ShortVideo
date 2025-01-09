package profile

import (
	"context"

	"github.com/Yamon955/ShortVideo/comm/base"
	"github.com/Yamon955/ShortVideo/comm/utils"
	"github.com/Yamon955/ShortVideo/protocol/user/pb"
	"github.com/Yamon955/ShortVideo/user/entity/def"
	"github.com/Yamon955/ShortVideo/user/entity/errcode"
	"github.com/Yamon955/ShortVideo/user/repo/mysql"
	MySQL "github.com/go-sql-driver/mysql"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/errs"
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
	rsp.FailedUids = make(map[uint64]string)
	// 去重
	req.Uids = utils.Duplicate(req.Uids)
	// 默认拉取个人主页资料
	if len(req.GetProfileTypes()) == 0 {
		req.ProfileTypes = append(req.GetProfileTypes(), pb.PROFILE_TYPES_MAIN_PAGE_INFO)
	}
	for _, uid := range req.GetUids() {
		if uid < def.MIN_UID {
			rsp.FailedUids[uid] = "uid is invalid"
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
					return
				})
			case pb.PROFILE_TYPES_LIKED_LIST_COUNT:
				handlers = append(handlers, func() (err error) {
					likedListCount, err = h.db.GetLikedListCount(ctx, uid)
					return
				})
			case pb.PROFILE_TYPES_COLLECT_LIST_COUNT:
				handlers = append(handlers, func() (err error) {
					collectListCount, err = h.db.GetCollectListCount(ctx, uid)
					return
				})
			}
		}
		// 并行拉取
		if err := trpc.GoAndWait(handlers...); err != nil {
			log.ErrorContextf(ctx, "GoAndWait failed, uid:%d, err:%v", uid, err)
			rsp.FailedUids[uid] = err.Error()
			continue
		}
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
	user, _ := h.db.FindUserByUID(ctx, req.GetUid())
	for _, infoType := range req.GetProfileTypes() {
		switch infoType {
		case pb.PROFILE_TYPES_USERNAME:
			if err := checkUsername(req.GetUsername()); err != nil {
				return err
			}
			if user.Username != req.GetUsername() {
				updateInfo["username"] = req.GetUsername()
			}
		case pb.PROFILE_TYPES_AVATOR:
			if user.Avator != req.GetAvator() {
				updateInfo["avatar"] = req.GetAvator()
			}
		case pb.PROFILE_TYPES_SIGN:
			if user.Sign != req.GetSign() {
				updateInfo["sign"] = req.GetSign()
			}
		case pb.PROFILE_TYPES_GENDER:
			if user.Gender != req.GetGender() {
				updateInfo["gender"] = req.GetGender()
			}
		}
	}
	if len(updateInfo) == 0 {
		return nil
	}
	// 更新用户信息
	if err := h.db.UpdateUserInfo(ctx, req.GetUid(), updateInfo); err != nil {
		log.ErrorContextf(ctx, "UpdateUserInfo failed, uid:%d, err:%v", req.GetUid(), err)
		if mysqlErr, ok := err.(*MySQL.MySQLError); ok && mysqlErr.Number == def.MySQLErrCode_UsernameIsDuplicate {
			return errs.New(errcode.ErrUsernameIsUsed, "用户名被占用")
		}
		return errs.New(errcode.ErrUpdateUserInfo, "更新失败，请稍后重试！")
	}
	return nil
}

func fillUserInfo(user *base.User, publishListCount int64, likedListCount int64, collectListCount int64) *pb.UserInfo {
	userInfo := &pb.UserInfo{}
	if user != nil {
		userInfo = &pb.UserInfo{
			MainPageInfo: &pb.MainPageInfo{
				ID:       user.ID,
				Username: user.Username,
				Avator:   user.Avator,
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
func checkUsername(username string) error {
	if len(username) == 0 || len(username) > def.MAX_LEN {
		return errs.New(errcode.ErrUsername, "用户名不能为空或者超过24个字符(8个汉字)")
	}
	return nil
}
