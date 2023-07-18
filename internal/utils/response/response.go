package response

import (
	"ccps.com/internal/utils"
	"ccps.com/internal/utils/language"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"net/http"
)

type SendHandler func(w http.ResponseWriter, data any, next func())
type ErrorHandler func(w http.ResponseWriter, err error, next func())
type JsonResponse struct {
	statusCode   int
	writer       http.ResponseWriter
	sendHandler  SendHandler
	langLoader   *language.Loader
	data         any
	errorHandler ErrorHandler
}

func NewJson(w http.ResponseWriter) *JsonResponse {
	langLoader := language.Load(map[string]map[string]string{})
	return &JsonResponse{langLoader: langLoader, writer: w}
}

func (r *JsonResponse) R(data any) {
	r.SetData(data).Send()
}

func (r *JsonResponse) E(errCode int, errMsg string) {
	r.Error(NewApiErr(http.StatusInternalServerError, errMsg, errCode))
}

func (r *JsonResponse) K() {
	r.data = utils.Nil(r.data, any(map[string]string{}))
	r.Send()
}

func (r *JsonResponse) SetLangLoader(loader *language.Loader) *JsonResponse {
	r.langLoader = loader
	return r
}

func (r *JsonResponse) SetSendHandler(handler SendHandler) *JsonResponse {
	r.sendHandler = handler
	return r
}

func (r *JsonResponse) SetData(data any) *JsonResponse {
	r.data = data
	return r
}

func (r *JsonResponse) GetData() any {
	return r.data
}

func (r *JsonResponse) SetStatusCode(code int) *JsonResponse {
	r.statusCode = code
	return r
}

func (r *JsonResponse) GetStatusCode() int {
	return r.statusCode
}

func (r *JsonResponse) SetErrorHandler(handler ErrorHandler) *JsonResponse {
	r.errorHandler = handler
	return r
}

func (r *JsonResponse) Error(err error) {
	var data ApiErr
	data.StatusCode = http.StatusInternalServerError
	data.ErrCode = CodeServerException
	data.ErrMsg = r.langLoader.Trans(MsgServerInternalError).
		Default("The network service request failed. Please try again later")

	switch err {
	case sqlx.ErrNotFound:
		data.StatusCode = http.StatusNotFound
		data.ErrCode = CodeDbNotFound
		data.ErrMsg = r.langLoader.Trans(MsgDbNotFound).
			Default("The data does not exist or has been deleted")
		break
	case ErrHttpNotFound:
		data.StatusCode = http.StatusNotFound
		data.ErrCode = http.StatusNotFound
		data.ErrMsg = r.langLoader.Trans(MsgHttpNotFound).
			Default("The specified resource does not exist")
		break
	case ErrHttpForbidden:
		data.StatusCode = http.StatusForbidden
		data.ErrCode = http.StatusForbidden
		data.ErrMsg = r.langLoader.Trans(MsgHttpForbidden).
			Default("You do not have permission to access the specified resource")
		break
	case ErrHttpUnauthorized:
		data.StatusCode = http.StatusUnauthorized
		data.ErrCode = http.StatusUnauthorized
		data.ErrMsg = r.langLoader.Trans(MsgHttpUnauthorized).
			Default("Unauthorized requests")
		break
	case ErrParamValidateFail:
		data.StatusCode = http.StatusUnprocessableEntity
		data.ErrCode = CodeParamValidateError
		data.ErrMsg = r.langLoader.Trans(MsgParamError).
			Default("Parameter input error, please check and modify before trying again")
		break
	case ErrSignatureExpiration:
		data.StatusCode = http.StatusForbidden
		data.ErrCode = CodeSignatureExpiration
		data.ErrMsg = r.langLoader.Trans(MsgSignatureExpiration).
			Default("Signature has expired")
		break
	case ErrSignatureInvalid:
		data.StatusCode = http.StatusForbidden
		data.ErrCode = CodeSignatureInvalid
		data.ErrMsg = r.langLoader.Trans(MsgSignatureInvalid).
			Default("Invalid signature")
		break
	case ErrTokenExpiration:
		data.StatusCode = http.StatusForbidden
		data.ErrCode = CodeTokenExpiration
		data.ErrMsg = r.langLoader.Trans(MsgTokenExpiration).
			Default("Token has expired")
		break
	case ErrTokenInvalid:
		data.StatusCode = http.StatusForbidden
		data.ErrCode = CodeTokenInvalid
		data.ErrMsg = r.langLoader.Trans(MsgTokenInvalid).
			Default("Invalid token")
		break
	case ErrServerException:
		data.StatusCode = http.StatusInternalServerError
		data.ErrCode = CodeServerException
		data.ErrMsg = r.langLoader.Trans(MsgServerInternalError).
			Default("The network service request failed. Please try again later")
		break
	case ErrHttpBadRequest:
		data.StatusCode = http.StatusBadRequest
		data.ErrCode = http.StatusBadRequest
		data.ErrMsg = err.Error()
		break
	case ErrAccountNotExist:
		data.StatusCode = http.StatusUnauthorized
		data.ErrCode = CodeAccountNotExist
		data.ErrMsg = r.langLoader.Trans(MsgAccountNotExist).
			Default("Account input error")
		break
	case ErrAccountPwdError:
		data.StatusCode = http.StatusUnauthorized
		data.ErrCode = CodeAccountPwdError
		data.ErrMsg = r.langLoader.Trans(MsgAccountPwdError).
			Default("Password input error")
		break
	}

	switch err.(type) {
	case *ApiErr:
		_e := err.(*ApiErr)
		data.StatusCode = _e.StatusCode
		data.ErrCode = _e.ErrCode
		data.ErrMsg = _e.ErrMsg
		break
	case validator.ValidationErrors:
		data.StatusCode = http.StatusUnprocessableEntity
		data.ErrCode = CodeParamValidateError
		data.ErrMsg = r.langLoader.
			TransByValidator(err.(validator.ValidationErrors)).
			Default("Parameter input error, please check and modify before trying again")
		break
	}

	next := func() {
		r.statusCode = utils.EmptyInt(r.statusCode, data.StatusCode)
		if r.data == nil {
			r.data = map[string]any{
				"error_code": data.ErrCode,
				"error_msg":  data.ErrMsg,
			}
		}

		r.Send()
	}

	if r.errorHandler != nil {
		r.errorHandler(r.writer, err, next)
		return
	}

	next()
}

func (r *JsonResponse) Send() {
	r.writer.Header().Add("Content-Type", "application/json")
	next := func() {
		r.writer.WriteHeader(utils.EmptyInt(r.statusCode, http.StatusOK))
		marshal, _ := json.Marshal(&r.data)
		_, _ = r.writer.Write(marshal)
	}

	if r.sendHandler != nil {
		r.sendHandler(r.writer, r.data, next)
		return
	}

	next()
}
