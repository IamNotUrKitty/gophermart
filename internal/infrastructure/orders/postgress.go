package orders

import (
	"context"

	"github.com/IamNotUrKitty/gophermart/internal/domain/order"
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

	if _, err := pool.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS orders (
		"number" bigint NOT NULL PRIMARY KEY,
		"status" status,
		"user_id" uuid REFERENCES users (id),
		"accrual" int,
		"uploaded_at" timestamp without time zone DEFAULT CURRENT_TIMESTAMP)`); err != nil {
		return nil, err
	}

	return &PostgressRepo{
		db: pool,
	}, nil
}

func (r *PostgressRepo) SaveOrder(ctx context.Context, o order.Order) error {
	_, err := r.db.Exec(ctx, "INSERT INTO orders (number, status, user_id, accrual) VALUES ($1, $2, $3)", o.Number(), o.Status(), o.UserID(), o.Accrual())
	if err != nil {
		return err
	}

	return err
}

func (r *PostgressRepo) GetOrder(ctx context.Context, number int) (*order.Order, error) {
	userRow := r.db.QueryRow(ctx, "SELECT number, status, user_id, accrual, uploaded_at FROM orders where number=$1", number)

	var o order.StoredOrder

	if err := userRow.Scan(&o.Number, &o.Status, &o.Accrual, &o.UploadedAt); err != nil {
		return nil, err
	}

	return order.NewOrder(), nil
}
