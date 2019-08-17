package views
const (
	ERROR_EXIST_TAG       = 10001
)
func ErrorRes(code int, err string) *Response {
	return &Response{
		Code:code,
		Error: err,
		Data:nil,
	}
}
