package repository

import (
	"context"

	"github.com/google/uuid"
)

type User struct {
	ID     uuid.UUID
	Name   string
	Email  string
	Status string
}

type PasswordCredential struct {
	UserId       uuid.UUID
	PasswordHash string
}

type Repository interface {
	CreateUser(
		ctx context.Context,
		user User,
		credential PasswordCredential,
	) error
}
