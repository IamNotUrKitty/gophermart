package orders

import (
	"context"

	"github.com/IamNotUrKitty/gophermart/internal/domain/order"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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

func (r *PostgressRepo) SaveOrder(ctx context.Context, o *order.Order, userID uuid.UUID) error {
	_, err := r.db.Exec(ctx, "INSERT INTO orders (number, user_id) VALUES ($1, $2)", o.Number, userID)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgressRepo) GetOrder(ctx context.Context, ord *order.Order) (*order.Order, error) {
	orderRow := r.db.QueryRow(ctx, "SELECT number, status, user_id, accrual, uploaded_at FROM orders where number=$1", ord.Number)

	var o order.Order

	if err := orderRow.Scan(&o.Number, &o.Status, &o.UserID, &o.Accrual, &o.UploadedAt); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &o, nil
}

func (r *PostgressRepo) GetOrdersByUserID(ctx context.Context, userID uuid.UUID) (*[]order.Order, error) {
	rows, err := r.db.Query(ctx, "SELECT number, status, user_id, accrual, uploaded_at FROM orders where user_id=$1", userID)
	if err != nil {
		return nil, err
	}

	orders, err := pgx.CollectRows(rows, pgx.RowToStructByName[order.Order])
	if err != nil {
		return nil, err
	}

	return &orders, nil
}
