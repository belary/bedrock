package utils

import "fmt"

type AppError struct {
	Code    string
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func NewAppError(code string, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
