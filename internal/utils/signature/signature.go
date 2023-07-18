package signature

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"net/url"
	"sort"
)

func Generate(secret string, str string) string {
	b := []byte(secret)
	mac := hmac.New(sha512.New, b)
	mac.Write([]byte(str))

	sign := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return base64.StdEncoding.EncodeToString([]byte(sign))
}

func GetSignStr(m map[string]string) string {
	var keys []string

	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	var queryParams = url.Values{}
	for _, key := range keys {
		queryParams.Add(key, m[key])
	}

	return queryParams.Encode()
}

func Equal(s1 string, s2 string) bool {
	return hmac.Equal([]byte(s1), []byte(s2))
}
