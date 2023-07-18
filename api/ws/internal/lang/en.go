package lang

import "ccps.com/internal/utils/response"

var EnLanguageMap = map[string]string{
	response.MsgHttpOk:              "OK",
	response.MsgHttpNotFound:        "Unable to find the specified resource",
	response.MsgHttpUnauthorized:    "Unauthenticated requests",
	response.MsgServerInternalError: "Internal Server Error",
	response.MsgParamError:          "Parameter validation error",
	response.MsgSignatureInvalid:    "Invalid signature",
	response.MsgSignatureExpiration: "Signature Expiration",
	response.MsgTokenInvalid:        "Invalid token",
	response.MsgTokenExpiration:     "Token expiration",

	"Validator.CreateClient.Name.required":   "Name is required",
	"Validator.CreateClient.RoomId.required": "Room number cannot be empty",
}
