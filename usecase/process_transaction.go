package usecase

import (
	"time"

	"github.com/eltonfonseca/codebank/domains"
	"github.com/eltonfonseca/codebank/dto"
	"github.com/eltonfonseca/codebank/interfaces"
)

type UseCaseTransaction struct {
	Repository interfaces.Repository
}

func NewUseCaseTransaction(repo interfaces.Repository) UseCaseTransaction {
	return UseCaseTransaction{
		Repository: repo,
	}
}

func (u UseCaseTransaction) ProcessTransaction(dto dto.Transaction) (domains.Transaction, error) {
	cc := newCC(dto)
	ccBalanceAndLimit, err := u.Repository.GetCreditCard(*cc)

	if err != nil {
		return domains.Transaction{}, err
	}

	cc.ID = ccBalanceAndLimit.ID
	cc.Limit = ccBalanceAndLimit.Limit
	cc.Balance = ccBalanceAndLimit.Balance

	t := newTransaction(dto, ccBalanceAndLimit)
	t.Process(cc)

	err = u.Repository.Save(*t, *cc)

	if err != nil {
		return domains.Transaction{}, err
	}

	return *t, nil
}

func newCC(dto dto.Transaction) *domains.CreditCard {
	cc := domains.NewCreditCard()
	cc.Name = dto.Name
	cc.Number = dto.Number
	cc.ExpirationMonth = dto.ExpirationMonth
	cc.ExpirationYear = dto.ExpirationYear
	cc.CVV = dto.CVV

	return cc
}

func newTransaction(dto dto.Transaction, cc domains.CreditCard) *domains.Transaction {
	t := domains.NewTransaction()
	t.CreditCardId = cc.ID
	t.Amount = dto.Amount
	t.Store = dto.Store
	t.Description = dto.Description
	t.CreatedAt = time.Now()

	return t
}
