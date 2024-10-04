package types

type Response struct {
	ResponseCode    string      `json:"response_code"`
	ResponseMessage string      `json:"response_message"`
	ResponseData    interface{} `json:"response_data,omitempty"`
}

type CustomError struct {
	Code    string
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func Error(code string, message string) error {
	return &CustomError{Code: code, Message: message}
}
