package domains

import (
	"time"

	"github.com/google/uuid"
)

type CreditCard struct {
	ID              uuid.UUID
	Name            string
	Number          string
	ExpirationMonth int32
	ExpirationYear  int32
	CVV             int32
	Balance         float64
	Limit           float64
	CreatedAt       time.Time
}

func NewCreditCard() *CreditCard {
	c := CreditCard{}
	c.ID = uuid.New()
	c.CreatedAt = time.Now()

	return &c
}
