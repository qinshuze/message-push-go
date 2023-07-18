package lang

import "ccps.com/internal/utils/response"

var CnLanguageMap = map[string]string{
	response.MsgHttpOk:              "成功",
	response.MsgHttpNotFound:        "找不到指定资源",
	response.MsgHttpUnauthorized:    "未认证的请求",
	response.MsgServerInternalError: "内部服务器错误",
	response.MsgParamError:          "参数验证错误",
	response.MsgSignatureInvalid:    "签名无效",
	response.MsgSignatureExpiration: "签名过期",
	response.MsgTokenInvalid:        "无效的令牌",
	response.MsgTokenExpiration:     "令牌过期",
	response.MsgDbNotFound:          "数据不存在或已被删除",
	response.MsgHttpForbidden:       "您没有访问指定资源的权限",

	AccountNotFound: "账号不存在",

	"Validator.TokenRequest.AccessKey.required": "access_key 不能为空",
	"Validator.TokenRequest.Signature.required": "签名不能为空",
	"Validator.TokenRequest.Name.required":      "名称不能为空",
	"Validator.TokenRequest.Expire.required":    "过期时间不能为空",
}
