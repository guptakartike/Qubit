package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/guptakartike/qubit/internal/auth"
)

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) CreateUser(
	ctx context.Context,
	user User,
	credential PasswordCredential,
) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(
		ctx,
		`INSERT INTO users (id, name, email, status)
		 VALUES ($1, $2, $3, $4)`,
		user.ID,
		user.Name,
		user.Email,
		user.Status,
	)
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) &&
			pgErr.Code == "23505" &&
			pgErr.ConstraintName == "users_email_active_idx" {
			return auth.ErrEmailAlreadyExists
		}

		return fmt.Errorf("create user: %w", err)
	}

	_, err = tx.Exec(
		ctx,
		`INSERT INTO password_credentials (user_id, password_hash)
		 VALUES ($1, $2)`,
		credential.UserId,
		credential.PasswordHash,
	)
	if err != nil {
		return fmt.Errorf("create password credential: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}
