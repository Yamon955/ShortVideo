// Package base 存放公共的基础功能包
package base

// User 用户个人主页信息
type User struct {
	ID           uint64 `gorm:"column:id;primaryKey;autoIncrement;index" json:"id"` // 自增主键，用户ID
	Username     string `gorm:"column:username;unique;index"  json:"username"`      // 用户名（唯一）
	Password     string `gorm:"column:password"  json:"password"`                   // 加密后的密码
	Avator       string `gorm:"column:avator" json:"avator"`                        // 简介
	Sign         string `gorm:"column:sign"  json:"sign"`                           // 头像
	Gender       int32  `gorm:"column:gender" json:"gender"`                        // 性别 0-未知 1-男 2-女
	FansCount    int32  `gorm:"column:fans_count" json:"fans_count"`                // 粉丝数
	FollowsCount int32  `gorm:"column:follows_count" json:"follows_count"`          // 关注数
	RegisterTime int64  `gorm:"column:register_time" json:"register_time"`          // 注册时间
}

// Video 视频信息
type Video struct {
	VID           uint64 `gorm:"column:vid;unique;index"  json:"vid"`         // 主键，视频ID
	UID           uint64 `gorm:"column:uid;index" json:"uid"`                 // 外键，视频发布者ID
	VideoURL      string `gorm:"column:video_url" json:"video_url"`           // 视频播放URL
	CoverURL      string `gorm:"column:cover_url" json:"cover_url"`           // 视频封面URL
	Title         string `gorm:"column:title" json:"title"`                   // 视频标题
	LikesCount    int32  `gorm:"column:likes_count" json:"likes_count"`       // 视频点赞数
	CollectsCount int32  `gorm:"column:collects_count" json:"collects_count"` // 视频收藏数
	CommentsCount int32  `gorm:"column:comments_count" json:"comments_count"` // 视频评论数
	PublishTime   int64  `gorm:"column:publish_time" json:"publish_time"`     // 发布时间
}

// Like 点赞信息
type Like struct {
	ID  uint64 `gorm:"column:id;primaryKey;autoIncrement;index" json:"id"` // 自增主键
	UID uint64 `gorm:"column:uid;index" json:"uid"`                        // 外键，点赞用户ID
	VID uint64 `gorm:"column:vid;index" json:"vid"`                        // 外键，被点赞视频ID
}

// Collect 收藏信息
type Collect struct {
	ID  uint64 `gorm:"column:id;primaryKey;autoIncrement;index" json:"id"` // 自增主键
	UID uint64 `gorm:"column:uid;index" json:"uid"`                        // 外键，收藏用户ID
	VID uint64 `gorm:"column:vid;index" json:"vid"`                        // 外键，被收藏视频ID
}
