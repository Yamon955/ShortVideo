package auth

import (
	"context"
	"time"

	"github.com/Yamon955/ShortVideo/comm/base"
	"github.com/Yamon955/ShortVideo/comm/utils"
	"github.com/Yamon955/ShortVideo/protocol/user/pb"
	"github.com/Yamon955/ShortVideo/user/entity/conf"
	"github.com/Yamon955/ShortVideo/user/entity/def"
	"github.com/Yamon955/ShortVideo/user/entity/errcode"
	"github.com/Yamon955/ShortVideo/user/repo/mysql"
	MySQL "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"trpc.group/trpc-go/trpc-go/errs"
	"trpc.group/trpc-go/trpc-go/log"
)

// handlerImpl 用户注册、登录处理器实现
type handlerImpl struct {
	db mysql.DBClient
}

// HandleRegister 用户注册
func (h *handlerImpl) HandleRegister(ctx context.Context, req *pb.RegisterReq) error {
	username := req.GetUsername()
	password := req.GetPassword()
	// 检查用户名和密码是否合法
	err := checkUsernameAndPassword(ctx, username, password)
	if err != nil {
		return err
	}
	// 新建新用户插入数据库
	if err := h.db.CreateUser(ctx, createUser(username, password)); err != nil {
		if mysqlErr, ok := err.(*MySQL.MySQLError); ok && mysqlErr.Number == def.MySQLErrCode_UsernameIsDuplicate {
			return errs.New(errcode.ErrUsernameIsUsed, "用户名已存在")
		}
		return errs.New(errcode.ErrDBOperation, "注册失败，请稍后重试！")
	}
	return nil
}

// HandleLogin 用户登录
func (h *handlerImpl) HandleLogin(ctx context.Context, req *pb.LoginReq, rsp *pb.LoginRsp) error {
	username := req.GetUsername()
	password := req.GetPassword()
	// 检查用户名和密码是否合法
	if err := checkUsernameAndPassword(ctx, username, password); err != nil {
		return err
	}
	user, err := h.db.FindUserByUsername(ctx, username)
	if err != nil {
		log.ErrorContextf(ctx, "FindUserByUsername failed, err:%v", err)
		if err == gorm.ErrRecordNotFound {
			return errs.New(errcode.ErrUserNotRegister, "用户名不存在")
		}
		return errs.New(errcode.ErrDBOperation, "登录失败，请稍后重试！")
	}
	if utils.Md5(req.Password) != user.Password {
		return errs.New(errcode.ErrPasswordNotMatch, "密码错误")
	}
	var token string
	token, err = utils.GenerateTokn(user.ID, username)
	if err != nil {
		log.ErrorContextf(ctx, "GenerateTokn failed, err:%v", err)
		return errs.New(errcode.ErrGenerateToekn, "登录失败，请稍后重试！")
	}
	rsp.Uid = user.ID
	rsp.Token = token
	return nil
}

// checkUsernameAndPassword 检查用户名和密码长度是否合法
func checkUsernameAndPassword(ctx context.Context, username string, password string) error {
	if len(username) == 0 || len(username) > def.MAX_LEN {
		log.ErrorContextf(ctx, "username[%s] is invalid", password)
		return errs.New(errcode.ErrUsername, "用户名不能为空或者超过24个字符(8个汉字)")
	}
	if len(password) < def.PWD_MIN_LEN || len(password) > def.MAX_LEN {
		log.ErrorContextf(ctx, "password[%s] is invalid", password)
		return errs.New(errcode.ErrPassword, "请输入8-24位密码")
	}
	return nil
}

// createUser 新建一个用户，使用默认配置初始化
func createUser(username string, pwd string) *base.User {
	return &base.User{
		Username:     username,
		Password:     utils.Md5(pwd),
		Avator:       conf.AppConf.UserDefaultConf.Avator,
		Sign:         conf.AppConf.UserDefaultConf.Sign,
		RegisterTime: time.Now().Unix(),
	}
}
