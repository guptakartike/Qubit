package auth

import (
	"net/mail"
	"strings"
	"unicode/utf8"
)

func ValidateRegisterRequest(req RegisterRequest) error {
	name := strings.TrimSpace(req.Name)
	email := strings.TrimSpace(req.Email)

	if name == "" {
		return &ValidationError{Field: "name", Message: "is required"}
	}

	if utf8.RuneCountInString(name) > 100 {
		return &ValidationError{Field: "name", Message: "must be 100 characters or fewer"}
	}

	if email == "" {
		return &ValidationError{Field: "email", Message: "is required"}
	}

	if len(email) > 255 {
		return &ValidationError{Field: "email", Message: "must be 255 characters or fewer"}
	}

	addr, err := mail.ParseAddress(email)
	if err != nil || addr.Name != "" {
		return &ValidationError{Field: "email", Message: "must be a valid email address"}
	}

	if req.Password == "" {
		return &ValidationError{Field: "password", Message: "is required"}
	}

	if utf8.RuneCountInString(req.Password) < 8 ||
		utf8.RuneCountInString(req.Password) > 127 {
		return &ValidationError{Field: "password", Message: "must be between 8 and 127 characters"}
	}

	return nil
}
