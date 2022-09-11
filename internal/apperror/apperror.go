package apperror

import (
	"errors"
	"fmt"
	"net/http"
)

type AppError struct {
	code    ErrCode
	message string
}

func New(code ErrCode, message string, args ...interface{}) AppError {
	return AppError{
		code:    code,
		message: fmt.Sprintf(message, args...),
	}
}

func (e AppError) Error() string {
	return e.message
}

func (e AppError) GetCode() ErrCode {
	return e.code
}

func IsAppError(err error, code ErrCode) bool {
	var appError AppError

	if errors.As(err, &appError) {
		return appError.GetCode() == code
	}
	return false
}

func GetStatusCode(err error, mapErrors map[ErrCode]int) int {
	var appError AppError

	if errors.As(err, &appError) {
		statusCode := mapErrors[appError.GetCode()]
		if statusCode != 0 {
			return statusCode
		}
	}

	return http.StatusInternalServerError
}
