package auth

import "github.com/google/uuid"

// User is the auth-domain representation of a registered user.
// It is returned by the service layer and consumed by handlers.
// It is not a database model.
type User struct {
	ID     uuid.UUID
	Name   string
	Email  string
	Status string
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}
