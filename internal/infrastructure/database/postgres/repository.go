package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/config"
)

type Repository struct {
	pool *pgxpool.Pool
}
type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	PasswordHash []byte `json:"password_hash"`
	RoleID       int    `json:"role_id"`
}

func NewUser(email string, role_ID int, passhash []byte) *User {
	return &User{
		ID:           uuid.New().String(),
		Email:        email,
		PasswordHash: passhash,
		RoleID:       role_ID,
	}
}
func NewRepo(ctx context.Context, config config.Database) *Repository {
	pg, err := New(ctx, config)
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	} else {
		InitTables(ctx, pg)
	}
	return &Repository{
		pool: pg,
	}
}
func (r *Repository) CloseConn() {
	if r.pool != nil {
		r.pool.Close()
	}
}
func (r *Repository) CreateUser(ctx context.Context, user User) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)
	query := `INSERT INTO users (id,email, password_hash, role_id) VALUES ($1, $2, $3,$4) ON CONFLICT (email) DO NOTHING`
	_, err = tx.Exec(ctx, query, user.ID, user.Email, user.PasswordHash, user.RoleID)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)
	query := `SELECT id, password_hash, role_id FROM users WHERE email = $1`
	var User User
	err = tx.QueryRow(ctx, query, email).Scan(&User.ID, &User.PasswordHash, &User.RoleID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return &User, nil
		}
		return &User, err
	}
	User.Email = email
	return &User, nil
}
