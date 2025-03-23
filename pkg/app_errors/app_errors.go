package app_errors

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(err error, status int) gin.H {
	return gin.H{
		"status":  status,
		"message": err.Error(),
	}
}

// ErrorType is the type of application error
type ErrorType string

const (
	// ErrorTypeValidation represents validation errors
	ErrorTypeValidation ErrorType = "validation"

	// ErrorTypeNotFound represents resource not found errors
	ErrorTypeNotFound ErrorType = "not_found"

	// ErrorTypeConflict represents resource conflict errors
	ErrorTypeConflict ErrorType = "conflict"

	// ErrorTypeUnauthorized represents authentication errors
	ErrorTypeUnauthorized ErrorType = "unauthorized"

	// ErrorTypeForbidden represents authorization errors
	ErrorTypeForbidden ErrorType = "forbidden"

	// ErrorTypeInternal represents internal server errors
	ErrorTypeInternal ErrorType = "internal"
)

// Common application errors
var (
	ErrInvalidRequest  = errors.New("invalid request")
	ErrInternalServer  = errors.New("internal server error")
	ErrNotFound        = errors.New("resource not found")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrForbidden       = errors.New("forbidden")
	ErrConflict        = errors.New("resource conflict")
	ErrInvalidToken    = errors.New("invalid token")
	ErrExpiredToken    = errors.New("token has expired")
	ErrInvalidOTP      = errors.New("invalid OTP")
	ErrExpiredOTP      = errors.New("OTP has expired")
	ErrUserExists      = errors.New("user already exists")
	ErrInvalidPassword = errors.New("invalid password")
)

// AppError represents an application error with type and optional details
type AppError struct {
	Err       error
	ErrorType ErrorType
	Details   string
}

// Error returns the error message
func (e *AppError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("%s: %s", e.Err.Error(), e.Details)
	}
	return e.Err.Error()
}

// Unwrap returns the wrapped error
func (e *AppError) Unwrap() error {
	return e.Err
}

// New creates a new AppError
func New(err error, errorType ErrorType, details string) *AppError {
	return &AppError{
		Err:       err,
		ErrorType: errorType,
		Details:   details,
	}
}

// NewValidationError creates a new validation error
func NewValidationError(details string) *AppError {
	return New(ErrInvalidRequest, ErrorTypeValidation, details)
}

// NewNotFoundError creates a new not found error
func NewNotFoundError(details string) *AppError {
	return New(ErrNotFound, ErrorTypeNotFound, details)
}

// NewConflictError creates a new conflict error
func NewConflictError(details string) *AppError {
	return New(ErrConflict, ErrorTypeConflict, details)
}

// NewUnauthorizedError creates a new unauthorized error
func NewUnauthorizedError(details string) *AppError {
	return New(ErrUnauthorized, ErrorTypeUnauthorized, details)
}

// NewForbiddenError creates a new forbidden error
func NewForbiddenError(details string) *AppError {
	return New(ErrForbidden, ErrorTypeForbidden, details)
}

// NewInternalError creates a new internal server error
func NewInternalError(details string) *AppError {
	return New(ErrInternalServer, ErrorTypeInternal, details)
}

// IsAppError checks if the error is an AppError
func IsAppError(err error) bool {
	var appErr *AppError
	return errors.As(err, &appErr)
}

// ErrorToStatusCode maps an AppError to an HTTP status code
func ErrorToStatusCode(err error) int {
	if !IsAppError(err) {
		return 500 // Internal Server Error
	}

	var appErr *AppError
	errors.As(err, &appErr)

	switch appErr.ErrorType {
	case ErrorTypeValidation:
		return 400 // Bad Request
	case ErrorTypeNotFound:
		return 404 // Not Found
	case ErrorTypeConflict:
		return 409 // Conflict
	case ErrorTypeUnauthorized:
		return 401 // Unauthorized
	case ErrorTypeForbidden:
		return 403 // Forbidden
	case ErrorTypeInternal:
		return 500 // Internal Server Error
	default:
		return 500 // Internal Server Error
	}
}
