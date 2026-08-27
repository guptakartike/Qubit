package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/guptakartike/qubit/internal/auth"
	"github.com/guptakartike/qubit/internal/auth/repository"
)

type PasswordHasher interface {
	HashPassword(password string) (string, error)
}

type RegistrationService struct {
	repo   repository.Repository
	hasher PasswordHasher
}

func NewRegistrationService(
	repo repository.Repository,
	hasher PasswordHasher,
) *RegistrationService {
	return &RegistrationService{
		repo:   repo,
		hasher: hasher,
	}
}

func (s *RegistrationService) Register(
	ctx context.Context,
	req auth.RegisterRequest,
) (auth.User, error) {
	if err := auth.ValidateRegisterRequest(req); err != nil {
		return auth.User{}, err
	}

	email := auth.NormaliseEmail(req.Email)

	passwordHash, err := s.hasher.HashPassword(req.Password)
	if err != nil {
		return auth.User{}, fmt.Errorf("hash password: %w", err)
	}

	userID := uuid.New()

	repoUser := repository.User{
		ID:     userID,
		Name:   strings.TrimSpace(req.Name),
		Email:  email,
		Status: "active",
	}

	credential := repository.PasswordCredential{
		UserId:       userID,
		PasswordHash: passwordHash,
	}

	if err := s.repo.CreateUser(ctx, repoUser, credential); err != nil {
		return auth.User{}, fmt.Errorf("create user: %w", err)
	}

	return auth.User{
		ID:     repoUser.ID,
		Name:   repoUser.Name,
		Email:  repoUser.Email,
		Status: repoUser.Status,
	}, nil
}
