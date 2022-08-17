package dto

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID              uuid.UUID
	Name            string
	Number          string
	ExpirationMonth int32
	ExpirationYear  int32
	CVV             int32
	Amount          float64
	Store           string
	Description     string
	CreatedAt       time.Time
}
