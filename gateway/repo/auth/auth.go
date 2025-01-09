package auth

import (
	"context"

	"github.com/Yamon955/ShortVideo/comm/utils"
	"trpc.group/trpc-go/trpc-go/log"
)

// VerifyJWT 校验登录态 token
func VerifyJWT(ctx context.Context, tokenStr string) (bool, uint64, error) {
	myClaims, err := utils.ParseToken(tokenStr)
	if err != nil {
		log.ErrorContextf(ctx, "ParseToken failed, err:%v", err)
		return false, 0, err
	}
	loginUid := myClaims.UID
	log.InfoContextf(ctx, "VerifyJWT success, uid:[%d]", loginUid)
	return true, loginUid, nil
}
