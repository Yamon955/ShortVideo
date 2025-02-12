package def

const (
	SUCCESS = 0

	MAX_LEN     = 24 // 用户名和密码最大长度
	PWD_MIN_LEN = 8  // 密码最低位数

	MIN_UID = 1 // 最小的合法 uid

	MySQLErrCode_UsernameIsDuplicate = 1062 // mysql 唯一键约束

	UserTagKeyPrefix      = "TAG:UID:"       // 存储 uid 的 tag 值的 key 前缀
	BloomFilter1Keysuffix = "_BloomFilter_1" // 用户布隆过滤器1 key 后缀
	BloomFilter2Keysuffix = "_BloomFilter_2" // 用户布隆过滤器2 key 后缀
	BloomFilterErrRate    = 0.01             // 布隆过滤器误差率
	BloomFilterCapacity   = 1000             // 布隆过滤器容量
)
