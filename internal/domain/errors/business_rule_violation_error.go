package apperrors

import "errors"

var (
	ErrInvalidInput = errors.New("invalid input")
	ErrNotFound     = errors.New("not found")
	ErrConflict     = errors.New("conflict")
)

type BusinessRuleViolationError struct {
	message string
}

func NewBusinessRuleViolationError(message string) *BusinessRuleViolationError {
	return &BusinessRuleViolationError{
		message: message,
	}
}
