package common

import (
	"errors"
	"fmt"
	"net/http"
)

type AppError struct {
	Code    int
	Message string
	Err     error
	Details any
}

func (e *AppError) Error() string {
	if e == nil {
		return ""
	}
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Err
}

func NewAppError(code int, message string, err error, details any) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
		Details: details,
	}
}

func BadRequest(message string, err error, details any) *AppError {
	return NewAppError(http.StatusBadRequest, message, err, details)
}

func Unauthorized(message string, err error, details any) *AppError {
	return NewAppError(http.StatusUnauthorized, message, err, details)
}

func Forbidden(message string, err error, details any) *AppError {
	return NewAppError(http.StatusForbidden, message, err, details)
}

func NotFound(message string, err error, details any) *AppError {
	return NewAppError(http.StatusNotFound, message, err, details)
}

func conflict(message string, err error, details any) *AppError {
	return NewAppError(http.StatusConflict, message, err, details)
}

func Unprocessable(message string, err error, details any) *AppError {
	return NewAppError(http.StatusUnprocessableEntity, message, err, details)
}

func Internal(message string, err error, details any) *AppError {
	if err != nil {
		err = errors.New(message + ": " + err.Error())
	}
	return NewAppError(http.StatusInternalServerError, message, err, details)
}

func AsAppError(err error) *AppError {
	if err == nil {
		return nil
	}
	var appErr *AppError
	if ok := errors.As(err, &appErr); ok {
		return appErr
	}
	return Internal("An unexpected error occurred", err, nil)
}

func Wwrapf(err error, format string, args ...any) *AppError {
	if err == nil {
		return nil
	}
	message := fmt.Sprintf(format, args...)
	return Internal("%s: %w", message, err, nil)
}
