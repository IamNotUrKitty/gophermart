package user

import (
	"context"

	"github.com/IamNotUrKitty/gophermart/internal/domain/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgressRepo struct {
	db *pgxpool.Pool
}

func NewPostgressRepo(cs string) (*PostgressRepo, error) {
	pool, err := pgxpool.New(context.Background(), cs)
	if err != nil {
		return nil, err
	}

	if _, err := pool.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS users (
		"id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
		"username" VARCHAR(250) NOT NULL UNIQUE,
		"password" VARCHAR(250))`); err != nil {
		return nil, err
	}

	return &PostgressRepo{
		db: pool,
	}, nil
}

func (r *PostgressRepo) SaveUser(ctx context.Context, u user.User) error {
	_, err := r.db.Exec(ctx, "INSERT INTO users (username, password) VALUES ($1, $2)", u.Username(), u.PasswordHash())
	if err != nil {
		return err
	}

	return err
}

func (r *PostgressRepo) GetUser(ctx context.Context, username string) (*user.User, error) {
	userRow := r.db.QueryRow(ctx, "SELECT id, username, password FROM users where username=$1", username)

	var u user.StoredUser

	if err := userRow.Scan(&u.ID, &u.Username, &u.Password); err != nil {
		return nil, err
	}

	return user.NewUser(u.ID, u.Username, u.Password), nil
}
