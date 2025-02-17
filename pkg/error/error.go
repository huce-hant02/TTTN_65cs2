package error

import "fmt"

type Error struct {
	sttCode int
	message string
	errCode string
}

func NewError(sttCode int, message string, errCode string) *Error {
	return &Error{
		sttCode: sttCode,
		message: message,
		errCode: errCode,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("[%d] - (%s) %s", e.sttCode, e.errCode, e.message)
}

func (e *Error) Message() string {
	return e.message

}

func (e *Error) ErrCode() string {
	return e.errCode
}

func (e *Error) SttCode() int {
	return e.sttCode
}
