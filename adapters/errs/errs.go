package errs

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) AppError {
	return AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewBadRequestError(message string) AppError {
	return AppError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewUnexpectedError(message string) AppError {
	return AppError{
		Code:    http.StatusExpectationFailed,
		Message: message,
	}
}
