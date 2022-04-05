package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"my_life/internal/domain"
)

type authPostgres struct {
	pool *pgxpool.Pool
}

func NewAuthPostgres(pool *pgxpool.Pool) *authPostgres {
	return &authPostgres{pool: pool}
}

const createUserQuery = `INSERT INTO users (username, phone, email, passwdHash, relevanceTime)
													Values ($1, $2, $3, $4, $5) RETURNING id;`

func (a authPostgres) CreateUser(ctx context.Context, u *domain.User) (int, error) {
	var userId int

	tx, err := a.pool.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("error initialising transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	row := a.pool.QueryRow(ctx, createUserQuery, u.Name, u.Phone, u.Email, u.Password, u.RelevanceTime)
	if err := row.Scan(&userId); err != nil {
		return 0, fmt.Errorf("error adding data to database: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf("error commiting transaction: %w", err)
	}

	return userId, nil
}

const getUserQuery = `SELECT id, username, phone, email, passwdHash, relevanceTime
								FROM users WHERE username = $1;`

func (a authPostgres) GetUser(ctx context.Context, username, password string) (*domain.User, error) {
	var u domain.User

	tx, err := a.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initialising transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	row := a.pool.QueryRow(ctx, getUserQuery, username)
	if err := row.Scan(&u.UId, &u.Name, &u.Phone, &u.Email, &u.Password, &u.RelevanceTime); err != nil {
		return nil, fmt.Errorf("error retreiveing data from database: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("error commiting transaction: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("wrong username or password")
	}

	return &u, nil
}
