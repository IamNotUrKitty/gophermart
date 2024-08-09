package handlers

import (
	"context"

	"github.com/IamNotUrKitty/gophermart/internal/domain/user"
)

type Repository interface {
	GetUser(ctx context.Context, userID string) (*user.User, error)
	SaveUser(ctx context.Context, u user.User) error
}

type Handler struct {
	address string
	repo    Repository
}

func NewHandler(address string, repo Repository) *Handler {
	return &Handler{
		address: address,
		repo:    repo,
	}
}
