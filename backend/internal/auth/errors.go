package auth

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidInput       = errors.New("invalid input")
	ErrEmailAlreadyExists = errors.New("email already exists")
)

// ValidationError is returned by ValidateRegisterRequest when input is invalid.
// It carries the field name and a human-readable message separately so that
// HTTP handlers can build structured JSON responses.
// errors.Is(err, ErrInvalidInput) returns true for any *ValidationError.
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func (e *ValidationError) Unwrap() error {
	return ErrInvalidInput
}
