package order

import "time"

type Status string

const (
	NEW        Status = "NEW"
	PROCESSING Status = "PROCESSING"
	INVALID    Status = "INVALID"
	PROCESSED  Status = "PROCESSED"
)

type Order struct {
	Number     string    `db:"number" json:"number"`
	Status     Status    `db:"status" json:"status"`
	Accrual    *float64  `db:"accrual" json:"accrual,omitempty"`
	UploadedAt time.Time `db:"uploaded_at" json:"uploaded_at"`
}

func NewOrder(number string) *Order {
	return &Order{
		Number: number,
	}
}

func CreateOrder(number string) (*Order, error) {
	return NewOrder(number), nil
}
