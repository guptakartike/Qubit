package service

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/guptakartike/qubit/internal/auth"
	"github.com/guptakartike/qubit/internal/auth/repository"
)

type fakeRepository struct {
	createUserCalled bool
	user             repository.User
	credential       repository.PasswordCredential
	err              error
}

func (f *fakeRepository) CreateUser(
	ctx context.Context,
	user repository.User,
	credential repository.PasswordCredential,
) error {
	f.createUserCalled = true
	f.user = user
	f.credential = credential

	return f.err
}

type fakeHasher struct {
	hash string
	err  error
}

func (f *fakeHasher) HashPassword(password string) (string, error) {
	return f.hash, f.err
}

func TestRegistrationService_Register(t *testing.T) {
	repo := &fakeRepository{}

	hasher := &fakeHasher{
		hash: "fake-password-hash",
	}

	service := NewRegistrationService(repo, hasher)

	req := auth.RegisterRequest{
		Name:     "  Test User  ",
		Email:    " TEST@EXAMPLE.COM ",
		Password: "correct-password",
	}

	user, err := service.Register(context.Background(), req)

	if err != nil {
		t.Fatalf("Register() error = %v", err)
	}

	if !repo.createUserCalled {
		t.Fatal("repository CreateUser was not called")
	}

	if user.Name != "Test User" {
		t.Errorf("user.Name = %q, want %q", user.Name, "Test User")
	}

	if user.Email != "test@example.com" {
		t.Errorf("user.Email = %q, want %q", user.Email, "test@example.com")
	}

	if user.Status != "active" {
		t.Errorf("user.Status = %q, want %q", user.Status, "active")
	}

	if repo.credential.PasswordHash != "fake-password-hash" {
		t.Errorf(
			"password hash = %q, want %q",
			repo.credential.PasswordHash,
			"fake-password-hash",
		)
	}

	if repo.credential.UserId != user.ID {
		t.Error("credential UserID does not match user ID")
	}
}

func TestRegistrationService_Register_ValidationFailure(t *testing.T) {
	repo := &fakeRepository{}

	hasher := &fakeHasher{
		hash: "fake-password-hash",
	}

	service := NewRegistrationService(repo, hasher)

	req := auth.RegisterRequest{
		Name:     "",
		Email:    "test@example.com",
		Password: "correct-password",
	}

	_, err := service.Register(context.Background(), req)

	if err == nil {
		t.Fatal("Register() expected an error, got nil")
	}

	if repo.createUserCalled {
		t.Fatal("repository should not be called when validation fails")
	}
}

func TestRegistrationService_Register_HashingFailure(t *testing.T) {
	repo := &fakeRepository{}

	hashError := errors.New("hashing failed")

	hasher := &fakeHasher{
		err: hashError,
	}

	service := NewRegistrationService(repo, hasher)

	req := auth.RegisterRequest{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "correct-password",
	}

	_, err := service.Register(context.Background(), req)

	if err == nil {
		t.Fatal("Register() expected an error, got nil")
	}

	if repo.createUserCalled {

		t.Fatal("repository should not be called when password hashing fails")
	}

}

func TestRegistrationService_Register_RepositoryFailure(t *testing.T) {
	repoError := errors.New("database unavailable")

	repo := &fakeRepository{
		err: repoError,
	}

	hasher := &fakeHasher{
		hash: "fake-password-hash",
	}

	service := NewRegistrationService(repo, hasher)

	req := auth.RegisterRequest{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "correct-password",
	}

	_, err := service.Register(context.Background(), req)

	if err == nil {
		t.Fatal("Register() expected an error, got nil")
	}

	if !repo.createUserCalled {
		t.Fatal("repository should have been called")
	}
}

func TestRegistrationService_Register_GeneratesUniqueIDs(t *testing.T) {
	repo := &fakeRepository{}

	hasher := &fakeHasher{
		hash: "fake-password-hash",
	}

	service := NewRegistrationService(repo, hasher)

	req := auth.RegisterRequest{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "correct-password",
	}

	user1, err := service.Register(context.Background(), req)
	if err != nil {
		t.Fatalf("first Register() error = %v", err)
	}

	user2, err := service.Register(context.Background(), req)
	if err != nil {
		t.Fatalf("second Register() error = %v", err)
	}

	if user1.ID == uuid.Nil {
		t.Fatal("first user ID was not generated")
	}

	if user2.ID == uuid.Nil {
		t.Fatal("second user ID was not generated")
	}

	if user1.ID == user2.ID {
		t.Fatal("two registrations generated the same user ID")
	}
}
