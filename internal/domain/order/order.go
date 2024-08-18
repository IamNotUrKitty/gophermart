package order

import "github.com/google/uuid"

type Status string

const (
	NEW        Status = "NEW"
	PROCESSING Status = "PROCESSING"
	INVALID    Status = "INVALID"
	PROCESSED  Status = "PROCESSED"
	REGISTERED Status = "REGISTERED"
)

type Order struct {
	number     int
	status     Status
	userID     uuid.UUID
	accrual    int
	uploadedAt string
}

type StoredOrder struct {
	Number     int       `json:"number"`
	Status     Status    `json:"status"`
	UserID     uuid.UUID `json:"userId"`
	Accrual    int       `json:"accrual"`
	UploadedAt string    `json:"uploaded_at"`
}

func NewOrder() *Order {
	return &Order{}
}

func CreateOrder() (*Order, error) {
	return NewOrder(), nil
}

func (o *Order) Number() int {
	return o.number
}

func (o *Order) Status() Status {
	return o.status
}

func (o *Order) UserID() uuid.UUID {
	return o.userID
}

func (o *Order) Accrual() int {
	return o.accrual
}

func (o *Order) UploadedAt() string {
	return o.uploadedAt
}
