package errcode

const (
	ErrReadConfig       = 20001 // 读取配置出错
	ErrUnmarshalConfig  = 20002 // 解析配置出错
	ErrUsername         = 20003 // 用户名有误
	ErrPassword         = 20004 // 密码有误
	ErrUsernameIsUsed   = 20005 // 用户名被占用
	ErrDBOperation      = 20006 // 数据库相关错误
	ErrUserNotRegister  = 20007 // 用户未注册
	ErrPasswordNotMatch = 20008 // 用户名密码不匹配
	ErrGenerateToekn    = 20009 // jwt 生成出错
	ErrInvalidToken     = 20010 // token 有误
	ErrUpdateUserInfo   = 20011 // 更新信息失败
)
