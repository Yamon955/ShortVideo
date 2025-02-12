package def

const (
	VideosInTwoMonthRedisKey = "videos_two_month"
	VideosCountEscape        = 20               // 每次拉取 20 个最近的视频
	BloomFilter1Keysuffix    = "_BloomFilter_1" // 用户布隆过滤器1 key 后缀
	BloomFilter2Keysuffix    = "_BloomFilter_2" // 用户布隆过滤器2 key 后缀
	VideoTagKeyPrefix        = "TAG:VID:"       // 存储 vid 的 tag 值的 key 前缀
	UserTagKeyPrefix         = "TAG:UID:"       // 存储 uid 的 tag 值的 key 前缀
	RecommendFeesCount       = 2                // 每次至少推荐 10 个视频
)
