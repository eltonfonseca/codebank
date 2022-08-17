package domains

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID           uuid.UUID
	Amount       float64
	Status       string
	Description  string
	Store        string
	CreditCardId uuid.UUID
	CreatedAt    time.Time
}

func NewTransaction() *Transaction {
	t := Transaction{}
	t.ID = uuid.New()
	t.CreatedAt = time.Now()

	return &t
}

func Validate(t *Transaction, cc CreditCard) {
	if (t.Amount + cc.Balance) > cc.Limit {
		t.Status = "rejected"
	} else {
		t.Status = "approved"
	}
}

func (t *Transaction) Process(cc *CreditCard) {
	Validate(t, *cc)

	if t.Status == "approved" {
		cc.Balance = cc.Balance + t.Amount
	}
}
