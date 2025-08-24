package utility

import (
	"Luckify/common/constants"
	"context"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

// GenerateJWT 生成一个 JWT
func GenerateJWT(secretKey string, iat, seconds, userId int64) (string, error) {
	// 创建令牌，设置声明
	claims := &jwt.MapClaims{
		"exp":                       iat + seconds,
		"iat":                       iat,
		constants.JwtClaimUserIdKey: userId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//使用密钥签名令牌
	return token.SignedString([]byte(secretKey))
}

// GetUserIdFromCtx 从上下文获取用户ID
func GetUserIdFromCtx(ctx context.Context) int64 {
	var uid int64
	if jsonUid, ok := ctx.Value(constants.JwtClaimUserIdKey).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			uid = int64Uid
		} else {
			logx.WithContext(ctx).Errorf("GetUidFromCtx err: %+v", err)
		}
	}
	return uid
}
