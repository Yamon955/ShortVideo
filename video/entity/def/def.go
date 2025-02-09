package def

const (
	DefaultPieceSize = 1024 * 1024 // 默认分片大小 1MB

	MinIOBucketName = "svtest" // MinIO 存储桶名

	MachineID = "machine-id" // Redis 记录初始化雪花算法机器 ID 的键名

	SQLLimitSize = 2 // 发布列表分页查询每次返回的数量
)
