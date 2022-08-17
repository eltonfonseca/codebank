package interfaces

import "github.com/eltonfonseca/codebank/domains"

type Repository interface {
	Save(t domains.Transaction, cc domains.CreditCard) error
	GetCreditCard(cc domains.CreditCard) (domains.CreditCard, error)
	CreateCreditCard(cc domains.CreditCard) error
}
