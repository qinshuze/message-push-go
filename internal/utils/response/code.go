package response

const (
	// CodeSignatureInvalid 1xxx 范围内的错误码归类为认证错误
	CodeSignatureInvalid    = 1001 // 签名无效
	CodeSignatureExpiration = 1002 // 签名过期
	CodeTokenInvalid        = 1003 // 令牌无效
	CodeTokenExpiration     = 1004 // 令牌过期
	CodeAccountNotExist     = 1005 // 账号不存在
	CodeAccountPwdError     = 1006 // 密码错误

	// CodeParamValidateError 2xxx 范围内的错误码归类为参数错误
	CodeParamValidateError = 2000 // 参数验证失败
	CodeAccountNotFound    = 2001 // 账号不存在

	// CodeDbNotFound 3xxx 范围内的错误码归类为数据库错误
	CodeDbNotFound = 3001 // 数据不存在

	// CodeRouteNotFound 4xxx 范围内的错误码归类为客户端错误，由客户端输入产生的错误
	CodeRouteNotFound = 4001 // 路由不存在

	// CodeServerException 5xxx 范围内的错误码归类为服务器错误
	CodeServerException = 5000 // 服务器异常
)

const (
	MsgHttpOk              = "Http.Ok"
	MsgHttpNotFound        = "Http.NotFound"
	MsgHttpUnauthorized    = "Http.Unauthorized"
	MsgHttpForbidden       = "Http.Forbidden"
	MsgServerInternalError = "Server.InternalError"
	MsgParamError          = "Param.Error"
	MsgSignatureInvalid    = "Signature.Invalid"
	MsgSignatureExpiration = "Signature.Expiration"
	MsgTokenInvalid        = "Token.Invalid"
	MsgTokenExpiration     = "Token.Expiration"
	MsgDbNotFound          = "DB.NotFound"
	MsgAccountNotExist     = "Account.NotExist"
	MsgAccountPwdError     = "Account.PasswordError"
)
