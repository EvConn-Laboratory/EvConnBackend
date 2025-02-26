package errors

import (
	"errors"
	"fmt"
)

type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

var (
	ErrInvalidCredentials = &AppError{Code: 401, Message: "Invalid credentials"}
	ErrUnauthorized       = &AppError{Code: 401, Message: "Unauthorized"}
	ErrForbidden          = &AppError{Code: 403, Message: "Forbidden"}
	ErrNotFound           = &AppError{Code: 404, Message: "Resource not found"}
	ErrValidation         = &AppError{Code: 422, Message: "Validation error"}
	ErrInternal           = &AppError{Code: 500, Message: "Internal server error"}
	ErrDatabaseOperation  = &AppError{Code: 500, Message: "Database operation failed"}
	ErrUserNotFound       = errors.New("user not found")
)

func NewError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
