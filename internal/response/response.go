package response

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func Error(message string, error interface{}) Response {
	return Response{
		Message: message,
		Error:   error,
	}
}

func Success(message string, data interface{}) Response {
	return Response{
		Message: message,
		Data:    data,
	}
}
