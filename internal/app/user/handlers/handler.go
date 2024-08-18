package handlers

import (
	"context"

	"github.com/IamNotUrKitty/gophermart/internal/domain/order"
	"github.com/IamNotUrKitty/gophermart/internal/domain/user"
)

type UserRepository interface {
	GetUser(ctx context.Context, userID string) (*user.User, error)
	SaveUser(ctx context.Context, u user.User) error
}

type OrderRepository interface {
	SaveOrder(ctx context.Context, o order.Order) error
	GetOrdersByUserId(ctx context.Context, userID string) ([]*order.Order, error)
}

type Handler struct {
	address  string
	userRepo UserRepository
}

func NewHandler(address string, repo UserRepository) *Handler {
	return &Handler{
		address:  address,
		userRepo: repo,
	}
}
