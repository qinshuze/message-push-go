package token

import (
	"ccps.com/internal/utils"
	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	Expire  int64
	Iat     int64
	Payload string
	jwt.RegisteredClaims
}

// Generate 生成Jwt令牌
// @secret: JWT 加解密密钥
// @seconds: 过期时间，单位秒
// @payload: 数据载体
func Generate(secret string, seconds int, payload string) (string, error) {
	iat := utils.NowUnixSecond()
	claims := Token{
		Expire:           iat + int64(seconds),
		Iat:              iat,
		Payload:          payload,
		RegisteredClaims: jwt.RegisteredClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// Parse 解析Jwt令牌
// @secret: JWT 加解密密钥
// @tokenStr: JWT令牌
func Parse(secret string, tokenStr string) (*Token, error) {
	token := Token{}
	_, err := jwt.ParseWithClaims(tokenStr, &token, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	return &token, err
}
