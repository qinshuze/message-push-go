package response

import "errors"

type ApiErr struct {
	StatusCode int
	ErrMsg     string
	ErrCode    int
}

func NewApiErr(statusCode int, errMsg string, errCode int) *ApiErr {
	return &ApiErr{statusCode, errMsg, errCode}
}

func (a *ApiErr) Error() string {
	return a.ErrMsg
}

var (
	ErrHttpNotFound        = errors.New("the resource does not exist or has been deleted")
	ErrHttpUnauthorized    = errors.New("unauthenticated requests")
	ErrHttpForbidden       = errors.New("forbidden requests")
	ErrServerException     = errors.New("internal server error")
	ErrHttpBadRequest      = errors.New("bad request")
	ErrParamValidateFail   = errors.New("parameter validation fail")
	ErrSignatureInvalid    = errors.New("invalid signature")
	ErrSignatureExpiration = errors.New("signature expiration")
	ErrTokenInvalid        = errors.New("invalid token")
	ErrTokenExpiration     = errors.New("token expiration")
	ErrAccountNotExist     = errors.New("account not exits")
	ErrAccountPwdError     = errors.New("password input error")
)
