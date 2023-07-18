package utils

import (
	"math"
	"time"
	"unsafe"
)

//// GenJwtToken 生成Jwt令牌
//// @secretKey: JWT 加解密密钥
//// @seconds: 过期时间，单位秒
//// @payload: 数据载体
//func GenJwtToken(secretKey string, seconds int64, payload string) (string, error) {
//	iat := time.Now().UnixMilli() / 1000
//	claims := make(jwt.MapClaims)
//	claims["exp"] = iat + seconds
//	claims["iat"] = iat
//	claims["payload"] = payload
//	token := jwt.New(jwt.SigningMethodHS256)
//	token.Claims = claims
//	return token.SignedString([]byte(secretKey))
//}
//
//func VerifyJwtToken(secretKey string, token string) {
//	_, _ = jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
//		log.Println(t.Claims)
//		return nil, nil
//	})
//}

type H map[string]any

func ShortExp(exp bool, a any, b any) any {
	if exp {
		return a
	} else {
		return b
	}
}

func Nil[T any](v T, defaultVal T) T {
	if (*[2]uintptr)(unsafe.Pointer(&v))[1] == 0 {
		return defaultVal
	} else {
		return v
	}

	//if reflect.ValueOf(v).IsNil() {
	//	return defaultVal
	//} else {
	//	return v
	//}
}

func EmptyString(v string, defaultVal string) string {
	return ShortExp(v == "", defaultVal, v).(string)
}

func EmptyInt(v int, defaultVal int) int {
	return ShortExp(v == 0, defaultVal, v).(int)
}

func ArrayItem[T any](a []T, i int, defaultVal T) T {
	if len(a) > i {
		return a[i]
	}

	return defaultVal
}

func NowUnixSecond() int64 {
	return NowUnixMilli() / 1000
}

func NowUnixMilli() int64 {
	return time.Now().UnixMilli()
}

func NowUnixMicro() int64 {
	return time.Now().UnixMicro()
}

func NowUnixNano() int64 {
	return time.Now().UnixNano()
}

func FloatEqual(v1 float64, v2 float64) bool {
	return math.Abs(v1-v2) < 0.00001
}

func FilterEmptyString[T any](a []T) []T {
	var _arr []T
	for _, a2 := range a {
		switch any(a2).(type) {
		case string:
			if any(a2) == "" {
				continue
			}
		}
		_arr = append(_arr, a2)
	}

	return _arr
}
