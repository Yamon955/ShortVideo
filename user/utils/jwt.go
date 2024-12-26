package utils

import (
	"time"

	"github.com/Yamon955/ShortVideo/user/entity/errcode"
	"github.com/dgrijalva/jwt-go"
	"trpc.group/trpc-go/trpc-go/errs"
)

const (
	SignKey           = "sv_auth_jwt" // 签名密钥
	DefaultExpireTime = 60            // 默认 token 过期时间 60min
)

type MyClaims struct {
	UID      uint64 `json:"uid"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateTokn 通过对称加密算法 HS256 生成 jwt
func GenerateTokn(uid uint64, username string) (string, error) {
	claims := MyClaims{
		UID:      uid,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "ShortVideo_Server",                                                   // 签发者
			Subject:   "yamon",                                                               // 签发对象
			Audience:  "sv_user",                                                             // 签发受众
			ExpiresAt: time.Now().Add(time.Duration(DefaultExpireTime) * time.Minute).Unix(), // 过期时间
			NotBefore: time.Now().Add(time.Second).Unix(),                                    // 最早使用时间
			IssuedAt:  time.Now().Unix(),                                                     // 签发时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(SignKey))
	return tokenStr, err
}

// ParseToken 解析 jwt
func ParseToken(tokenStr string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})
	if err != nil {
		return nil, errs.New(errcode.ErrInvalidToken, "invalid token")
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errs.New(errcode.ErrInvalidToken, "invalid token")
}
