package httpx

import "go-layout/utils"

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func ErrorResponse(message string, error interface{}) Response {
	return Response{
		Message: utils.CapitalizeFirst(message),
		Error:   error,
	}
}

func SuccessResponse(message string, data interface{}) Response {
	return Response{
		Message: utils.CapitalizeFirst(message),
		Data:    data,
	}
}
