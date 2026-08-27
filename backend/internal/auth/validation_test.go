package auth

import (
	"errors"
	"strings"
	"testing"
)

func TestValidateRegisterRequest_Valid(t *testing.T) {
	req := RegisterRequest{
		Name:     "Kartike Gupta",
		Email:    "kartike@example.com",
		Password: "correctpassword",
	}

	if err := ValidateRegisterRequest(req); err != nil {
		t.Fatalf("ValidateRegisterRequest() unexpected error = %v", err)
	}
}

func TestValidateRegisterRequest_EmptyName(t *testing.T) {
	req := RegisterRequest{
		Name:     "",
		Email:    "kartike@example.com",
		Password: "correctpassword",
	}

	err := ValidateRegisterRequest(req)
	if err == nil {
		t.Fatal("expected error for empty name, got nil")
	}

	if !errors.Is(err, ErrInvalidInput) {
		t.Errorf("error = %v, want errors.Is(err, ErrInvalidInput) = true", err)
	}

	var ve *ValidationError
	if !errors.As(err, &ve) || ve.Field != "name" {
		t.Errorf("expected ValidationError with Field=name, got %v", err)
	}
}

func TestValidateRegisterRequest_NameTooLong(t *testing.T) {
	req := RegisterRequest{
		Name:     strings.Repeat("a", 101),
		Email:    "kartike@example.com",
		Password: "correctpassword",
	}

	err := ValidateRegisterRequest(req)
	if err == nil {
		t.Fatal("expected error for name > 100 runes, got nil")
	}

	var ve *ValidationError
	if !errors.As(err, &ve) || ve.Field != "name" {
		t.Errorf("expected ValidationError with Field=name, got %v", err)
	}
}

func TestValidateRegisterRequest_EmptyEmail(t *testing.T) {
	req := RegisterRequest{
		Name:     "Kartike Gupta",
		Email:    "",
		Password: "correctpassword",
	}

	err := ValidateRegisterRequest(req)
	if err == nil {
		t.Fatal("expected error for empty email, got nil")
	}

	var ve *ValidationError
	if !errors.As(err, &ve) || ve.Field != "email" {
		t.Errorf("expected ValidationError with Field=email, got %v", err)
	}
}

func TestValidateRegisterRequest_EmailTooLong(t *testing.T) {
	// 247 'a' chars + "@test.com" (9 chars) = 256 bytes, exceeding the 255-byte limit.
	req := RegisterRequest{
		Name:     "Kartike Gupta",
		Email:    strings.Repeat("a", 247) + "@test.com",
		Password: "correctpassword",
	}

	err := ValidateRegisterRequest(req)
	if err == nil {
		t.Fatal("expected error for email > 255 bytes, got nil")
	}

	var ve *ValidationError
	if !errors.As(err, &ve) || ve.Field != "email" {
		t.Errorf("expected ValidationError with Field=email, got %v", err)
	}
}

func TestValidateRegisterRequest_MalformedEmail(t *testing.T) {
	req := RegisterRequest{
		Name:     "Kartike Gupta",
		Email:    "not-an-email",
		Password: "correctpassword",
	}

	err := ValidateRegisterRequest(req)
	if err == nil {
		t.Fatal("expected error for malformed email, got nil")
	}

	var ve *ValidationError
	if !errors.As(err, &ve) || ve.Field != "email" {
		t.Errorf("expected ValidationError with Field=email, got %v", err)
	}
}

func TestValidateRegisterRequest_DisplayNameEmail(t *testing.T) {
	// net/mail.ParseAddress accepts "John Doe" <john@example.com>.
	// Registration should reject this and require a plain mailbox address.
	req := RegisterRequest{
		Name:     "Kartike Gupta",
		Email:    `"John Doe" <john@example.com>`,
		Password: "correctpassword",
	}

	err := ValidateRegisterRequest(req)
	if err == nil {
		t.Fatal("expected error for display-name email format, got nil")
	}

	var ve *ValidationError
	if !errors.As(err, &ve) || ve.Field != "email" {
		t.Errorf("expected ValidationError with Field=email, got %v", err)
	}
}

func TestValidateRegisterRequest_EmptyPassword(t *testing.T) {
	req := RegisterRequest{
		Name:     "Kartike Gupta",
		Email:    "kartike@example.com",
		Password: "",
	}

	err := ValidateRegisterRequest(req)
	if err == nil {
		t.Fatal("expected error for empty password, got nil")
	}

	var ve *ValidationError
	if !errors.As(err, &ve) || ve.Field != "password" {
		t.Errorf("expected ValidationError with Field=password, got %v", err)
	}
}

func TestValidateRegisterRequest_PasswordTooShort(t *testing.T) {
	req := RegisterRequest{
		Name:     "Kartike Gupta",
		Email:    "kartike@example.com",
		Password: "short",
	}

	err := ValidateRegisterRequest(req)
	if err == nil {
		t.Fatal("expected error for password < 8 chars, got nil")
	}

	var ve *ValidationError
	if !errors.As(err, &ve) || ve.Field != "password" {
		t.Errorf("expected ValidationError with Field=password, got %v", err)
	}
}

func TestValidateRegisterRequest_PasswordTooLong(t *testing.T) {
	req := RegisterRequest{
		Name:     "Kartike Gupta",
		Email:    "kartike@example.com",
		Password: strings.Repeat("a", 128),
	}

	err := ValidateRegisterRequest(req)
	if err == nil {
		t.Fatal("expected error for password > 127 chars, got nil")
	}

	var ve *ValidationError
	if !errors.As(err, &ve) || ve.Field != "password" {
		t.Errorf("expected ValidationError with Field=password, got %v", err)
	}
}

func TestValidateRegisterRequest_PasswordAtMinBoundary(t *testing.T) {
	req := RegisterRequest{
		Name:     "Kartike Gupta",
		Email:    "kartike@example.com",
		Password: strings.Repeat("a", 8),
	}

	if err := ValidateRegisterRequest(req); err != nil {
		t.Fatalf("expected no error for 8-char password, got %v", err)
	}
}

func TestValidateRegisterRequest_PasswordAtMaxBoundary(t *testing.T) {
	req := RegisterRequest{
		Name:     "Kartike Gupta",
		Email:    "kartike@example.com",
		Password: strings.Repeat("a", 127),
	}

	if err := ValidateRegisterRequest(req); err != nil {
		t.Fatalf("expected no error for 127-char password, got %v", err)
	}
}

func TestValidateRegisterRequest_WhitespaceName(t *testing.T) {
	// A name consisting entirely of whitespace is effectively empty.
	req := RegisterRequest{
		Name:     "   ",
		Email:    "kartike@example.com",
		Password: "correctpassword",
	}

	err := ValidateRegisterRequest(req)
	if err == nil {
		t.Fatal("expected error for whitespace-only name, got nil")
	}

	var ve *ValidationError
	if !errors.As(err, &ve) || ve.Field != "name" {
		t.Errorf("expected ValidationError with Field=name, got %v", err)
	}
}
