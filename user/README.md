# User--用户模块

## 功能简介

用户模块主要包括四个功能：用户注册/登录、用户资料获取/修改

```protobuf
service User {
    rpc Register(RegisterReq) returns (RegisterRsp);
    rpc Login(LoginReq) returns (LoginRsp);
    rpc BatchGetProfile(BatchGetProfileReq) returns (BatchGetProfileRsp);
    rpc SetProfile(SetProfileReq) returns (SetProfileRsp);
}
```

用户登录之后会返回给前端一个JWT，之后的前端的HTTP请求都会带上JWT来鉴权，这里把JWT放到HTTP头部的`Authorization`字段中，而不是`Cookie`中，避免跨域问题。

##  user表

```go
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
```