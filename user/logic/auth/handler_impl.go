package auth

import (
	"context"

	"github.com/Yamon955/ShortVideo/user/entity/conf"
	"github.com/Yamon955/ShortVideo/user/entity/errcode"
	"github.com/Yamon955/ShortVideo/user/entity/model"
	"github.com/Yamon955/ShortVideo/user/pb"
	"github.com/Yamon955/ShortVideo/user/repo/mysql"
	"github.com/Yamon955/ShortVideo/user/utils"
	"gorm.io/gorm"
	"trpc.group/trpc-go/trpc-go/errs"
	"trpc.group/trpc-go/trpc-go/log"
)

type authHandlerandlerImpl struct {
	db mysql.DBClient
}

const (
	maxLen = 24
	minLen = 8
)

func (h *authHandlerandlerImpl) HandleRegister(ctx context.Context, req *pb.RegisterReq) error {
	username := req.GetUsername()
	password := req.GetPassword()
	// 检查用户名和密码是否合法
	err := checkUsernameAndPassword(username, password)
	if err != nil {
		return err
	}
	// 检查用户名是否被占用
	user, err := h.db.FindUserByUsername(ctx, username)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.ErrorContextf(ctx, "FindUserByUsername failed, err:%v", err)
		return errs.New(errcode.ErrDBOperation, "数据库查询出错，请稍后重试！")
	}
	if user.ID != 0 {
		return errs.New(errcode.ErrUsernameIsUsed, "用户名被占用")
	}
	// 新建新用户插入数据库
	createUser(username, password)
	return nil
}

func (h *authHandlerandlerImpl) HandleLogin(ctx context.Context, req *pb.LoginReq, rsp *pb.LoginRsp) error {
	username := req.GetUsername()
	password := req.GetPassword()
	// 检查用户名和密码是否合法
	err := checkUsernameAndPassword(username, password)
	if err != nil {
		return err
	}
	user, err := h.db.FindUserByUsername(ctx, username)
	if err != nil {
		log.ErrorContextf(ctx, "FindUserByUsername failed, err:%v", err)
		if err == gorm.ErrRecordNotFound {
			return errs.New(errcode.ErrUserNotRegister, "用户名不存在")
		}
		return errs.New(errcode.ErrDBOperation, "数据库查询出错，请稍后重试！")
	}
	if utils.Md5(req.Password) != user.Password {
		return errs.New(errcode.ErrPasswordNotMatch, "密码错误")
	}
	token, err := utils.GenerateTokn(user.ID, username)
	if err != nil {
		log.ErrorContextf(ctx, "GenerateTokn failed, err:%v", err)
		return errs.New(errcode.ErrGenerateToekn, "token生成出错，请重新登录")
	}
	rsp = &pb.LoginRsp{
		Uid:   user.ID,
		Token: token,
	}
	return nil
}

// checkUsernameAndPassword 检查用户名和密码长度是否合法
func checkUsernameAndPassword(username string, password string) error {
	if len(username) == 0 || len(username) > maxLen {
		return errs.New(errcode.ErrUsername, "用户名不能为空或者超过24个字符(8个汉字)")
	}
	if len(password) < minLen || len(password) > maxLen {
		return errs.New(errcode.ErrPassword, "请输入8-24位密码")
	}
	return nil
}

// createUser 新建一个用户，使用默认配置初始化
func createUser(username string, pwd string) *model.User {
	return &model.User{
		Username:     username,
		Password:     utils.Md5(pwd),
		Avatar:       conf.AppConf.Avator,
		Sign:         conf.AppConf.Sign,
		FansCount:    conf.AppConf.FansCount,
		FollowsCount: conf.AppConf.FollowsCount,
	}
}
