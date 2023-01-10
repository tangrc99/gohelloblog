package err_code

import (
	"fmt"
	"net/http"
)

// Error represents server error.
type Error struct {
	ECode    int      `json:"code"`    // error code
	EMsg     string   `json:"msg"`     // error description
	EDetails []string `json:"details"` // error detail info
}

// codes records all possible error. Users should only get error from this map
var codes = map[int]string{}

// registryError generate a new Error struct and put it to map codes
func registryError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("Error Code %d Exits", code))
	}
	codes[code] = msg
	return &Error{ECode: code, EMsg: msg}
}

func (e *Error) Code() int {
	return e.ECode
}

func (e *Error) Message() string {
	return e.EMsg
}

func (e *Error) Detail() []string {
	return e.EDetails
}

func (e *Error) WithoutDetails() *Error {
	err := *e
	return &err
}

func (e *Error) WithDetails(details ...string) *Error {
	err := *e
	for _, d := range details {
		err.EDetails = append(err.EDetails, d) // 追加 slice 元素的方式
	}
	return &err
}

func (e *Error) Info() string {
	return fmt.Sprintf("错误码：%d, 错误信息:：%s", e.Code(), e.Message())
}

func (e *Error) HttpCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case NotFound.Code():
		return http.StatusNotFound
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	case MethodNotAllowed.Code():
		return http.StatusMethodNotAllowed
	case PermissionDenied.Code():
		return http.StatusForbidden
	}

	return http.StatusInternalServerError
}
