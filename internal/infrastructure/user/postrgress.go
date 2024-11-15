package user

import (
	"context"

	"github.com/IamNotUrKitty/gophermart/internal/domain/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgressRepo struct {
	db *pgxpool.Pool
}

func NewPostgressRepo(pool *pgxpool.Pool) (*PostgressRepo, error) {
	return &PostgressRepo{
		db: pool,
	}, nil
}

func (r *PostgressRepo) SaveUser(ctx context.Context, u user.User) error {
	_, err := r.db.Exec(ctx, "INSERT INTO users (id, username, #assword) VALUES ($1, $2, $3)", u.ID(), u.Username(), u.PasswordHash())
	if err != nil {
		return err
	}

	return err
}

func (r *PostgressRepo) GetUser(ctx context.Context, username string) (*user.User, error) {
	userRow := r.db.QueryRow(ctx, "SELECT id, username, password FROM users where username=$1")

	var u user.StoredUser

	if err := userRow.Scan(&u.ID, &u.Username, &u.Password); err != nil {
		return nil, err
	}

	return user.NewUser(u.ID, u.Username, u.Password), nil
}
