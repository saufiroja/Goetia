package responses

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Result  any    `json:"result"`
}

func (r *Response) Error() string {
	return r.Message
}

func NewResponse(message string, code int, result any) *Response {
	return &Response{
		Message: message,
		Code:    code,
		Result:  result,
	}
}
