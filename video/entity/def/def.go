package def

const (
	DefaultPieceSize = 1024 * 1024 // 默认分片大小 1MB
	MinIOBucketName  = "svtest"    // MinIO 存储桶名
	SQLLimitSize     = 2           // 发布列表分页查询每次返回的数量

	MachineID                    = "machine-id" // Redis 记录初始化雪花算法机器 ID 的键名
	VideoTagKeyPrefix            = "TAG:VID:"   // 存储 vid 的 tag 值的 key 前缀
	UserTagKeyPrefix             = "TAG:UID:"   // 存储 uid 的 tag 值的 key 前缀
	VideosInTwoMonthRedisKey     = "videos_two_month"
	UserCurrentBloomFilterPrefix = "current:bloomfilter:" // 用户当前正在使用的布隆过滤器
	BloomFilter1Keysuffix        = "_BloomFilter_1"       // 用户布隆过滤器1 key 后缀
	BloomFilter2Keysuffix        = "_BloomFilter_2"       // 用户布隆过滤器2 key 后缀
	BloomFilterErrRate           = 0.01                   // 布隆过滤器误差率
	BloomFilterCapacity          = 100000                 // 布隆过滤器容量
)
