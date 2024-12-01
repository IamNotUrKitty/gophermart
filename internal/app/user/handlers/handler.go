package handlers

import (
	"context"

	"github.com/IamNotUrKitty/gophermart/internal/domain/order"
	"github.com/IamNotUrKitty/gophermart/internal/domain/user"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetUser(ctx context.Context, userID string) (*user.User, error)
	SaveUser(ctx context.Context, u user.User) error
}

type OrderRepository interface {
	SaveOrder(ctx context.Context, o *order.Order, userID uuid.UUID) error
	GetOrdersByUserID(ctx context.Context, userID uuid.UUID) (*[]order.Order, error)
	GetOrder(ctx context.Context, o *order.Order) (*order.Order, error)
}

type Handler struct {
	address   string
	userRepo  UserRepository
	orderRepo OrderRepository
}

func NewHandler(address string, userRepo UserRepository, orderRepo OrderRepository) *Handler {
	return &Handler{
		address:   address,
		userRepo:  userRepo,
		orderRepo: orderRepo,
	}
}
