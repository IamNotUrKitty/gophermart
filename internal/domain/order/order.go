package order

import (
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
)

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
	UserID     uuid.UUID
}

func NewOrder(number string) *Order {
	return &Order{
		Number: number,
	}
}

func CreateOrder(number string) (*Order, error) {
	valid := ValidateOrderNumber(number)
	if !valid {
		return nil, errors.New("invalid order")
	}

	return NewOrder(number), nil
}

func ValidateOrderNumber(number string) bool {

	i, err := strconv.ParseInt(number, 10, 64)

	if err != nil {
		return false
	}

	return (i%10+checksum(i/10))%10 == 0
}

func checksum(number int64) int64 {
	var luhn int64

	for i := 0; number > 0; i++ {
		cur := number % 10

		if i%2 == 0 { // even
			cur = cur * 2
			if cur > 9 {
				cur = cur%10 + cur/10
			}
		}

		luhn += cur
		number = number / 10
	}
	return luhn % 10
}
