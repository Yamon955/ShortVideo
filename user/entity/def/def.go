package def

const (
	SUCCESS = 0

	MAX_LEN     = 24 // 用户名和密码最大长度
	PWD_MIN_LEN = 8  // 密码最低位数

	MIN_UID = 1 // 最小的合法 uid

	MySQLErrCode_UsernameIsDuplicate = 1062 // mysql 唯一键约束
)
