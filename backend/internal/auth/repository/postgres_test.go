package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/guptakartike/qubit/internal/auth"
)

func TestPostgresRepository_CreateUser(t *testing.T) {
	databaseURL := os.Getenv("TEST_DATABASE_URL")

	if databaseURL == "" {
		t.Fatal("TEST_DATABASE_URL is not set")
	}

	ctx := context.Background()

	db, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		t.Fatalf("create database pool: %v", err)
	}
	defer db.Close()

	repo := NewPostgresRepository(db)

	userID := uuid.New()

	user := User{
		ID:     userID,
		Name:   "Test User",
		Email:  fmt.Sprintf("test-%s@example.com", uuid.NewString()),
		Status: "active",
	}

	credential := PasswordCredential{
		UserId:       userID,
		PasswordHash: "test-password-hash",
	}

	err = repo.CreateUser(ctx, user, credential)
	if err != nil {
		t.Fatalf("CreateUser() error = %v", err)
	}

	t.Cleanup(func() {
		_, _ = db.Exec(ctx, `DELETE FROM password_credentials WHERE user_id = $1`, userID)
		_, _ = db.Exec(ctx, `DELETE FROM users WHERE id = $1`, userID)
	})

	var storedName string
	var storedEmail string

	err = db.QueryRow(
		ctx,
		`SELECT name, email FROM users WHERE id = $1`,
		userID,
	).Scan(&storedName, &storedEmail)

	if err != nil {
		t.Fatalf("query created user: %v", err)
	}

	if storedName != user.Name {
		t.Errorf("stored name = %q, want %q", storedName, user.Name)
	}

	if storedEmail != user.Email {
		t.Errorf("stored email = %q, want %q", storedEmail, user.Email)
	}
}

func TestPostgresRepository_CreateUser_RollsBackOnCredentialFailure(t *testing.T) {
	databaseURL := os.Getenv("TEST_DATABASE_URL")

	if databaseURL == "" {
		t.Fatal("TEST_DATABASE_URL is not set")
	}

	ctx := context.Background()

	db, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		t.Fatalf("create database pool: %v", err)
	}
	defer db.Close()

	repo := NewPostgresRepository(db)

	userID := uuid.New()
	invalidUserID := uuid.New()

	user := User{
		ID:     userID,
		Name:   "Rollback User",
		Email:  "rollback@example.com",
		Status: "active",
	}

	credential := PasswordCredential{
		UserId:       invalidUserID,
		PasswordHash: "test-password-hash",
	}

	err = repo.CreateUser(ctx, user, credential)

	if err == nil {
		t.Fatal("CreateUser() expected an error, got nil")
	}

	var count int

	err = db.QueryRow(
		ctx,
		`SELECT COUNT(*) FROM users WHERE id = $1`,
		userID,
	).Scan(&count)

	if err != nil {
		t.Fatalf("query user after rollback: %v", err)
	}

	if count != 0 {
		t.Fatalf("user still exists after rollback")
	}
}

func TestPostgresRepository_CreateUser_DuplicateEmail(t *testing.T) {
	databaseURL := os.Getenv("TEST_DATABASE_URL")

	if databaseURL == "" {
		t.Fatal("TEST_DATABASE_URL is not set")
	}

	ctx := context.Background()

	db, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		t.Fatalf("create database pool: %v", err)
	}
	defer db.Close()

	repo := NewPostgresRepository(db)

	email := fmt.Sprintf("duplicate-%s@example.com", uuid.NewString())

	user1 := User{
		ID:     uuid.New(),
		Name:   "First User",
		Email:  email,
		Status: "active",
	}

	credential1 := PasswordCredential{
		UserId:       user1.ID,
		PasswordHash: "hash-1",
	}

	if err := repo.CreateUser(ctx, user1, credential1); err != nil {
		t.Fatalf("first CreateUser() error = %v", err)
	}

	defer func() {
		_, _ = db.Exec(
			ctx,
			`DELETE FROM password_credentials WHERE user_id = $1`,
			user1.ID,
		)

		_, _ = db.Exec(
			ctx,
			`DELETE FROM users WHERE id = $1`,
			user1.ID,
		)
	}()

	user2 := User{
		ID:     uuid.New(),
		Name:   "Second User",
		Email:  email,
		Status: "active",
	}

	credential2 := PasswordCredential{
		UserId:       user2.ID,
		PasswordHash: "hash-2",
	}

	err = repo.CreateUser(ctx, user2, credential2)

	if !errors.Is(err, auth.ErrEmailAlreadyExists) {
		t.Fatalf(
			"CreateUser() error = %v, want %v",
			err,
			auth.ErrEmailAlreadyExists,
		)
	}

	var count int

	err = db.QueryRow(
		ctx,
		`SELECT COUNT(*) FROM users WHERE email = $1`,
		email,
	).Scan(&count)

	if err != nil {
		t.Fatalf("query users: %v", err)
	}

	if count != 1 {
		t.Fatalf("user count = %d, want 1", count)
	}
}
