package model

type User struct {
	ID           uint64 `gorm:"column:id;primaryKey;autoIncrement;index" json:"id"`
	Username     string `gorm:"column:username;unique;index"  json:"username"`
	Password     string `gorm:"column:password"  json:"password"`
	Avatar       string `gorm:"column:avatar" json:"avatar"`
	Sign         string `gorm:"column:sign"  json:"sign"`
	Gender       int    `gorm:"column:gender" json:"gender"`
	FansCount    int    `gorm:"column:fans_count" json:"fans_count"`
	FollowsCount int    `gorm:"column:follows_count" json:"follows_count"`
	RegisterTime int64  `gorm:"column:register_time" json:"register_time"`
}
