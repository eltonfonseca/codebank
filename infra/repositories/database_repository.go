package repositories

import (
	"database/sql"
	"errors"

	"github.com/eltonfonseca/codebank/domains"
)

type DatabaseRepository struct {
	db *sql.DB
}

func NewDatabaseRepository(db *sql.DB) *DatabaseRepository {
	return &DatabaseRepository{db: db}
}

func (repo *DatabaseRepository) Save(t domains.Transaction, cc domains.CreditCard) error {
	query := `INSERT INTO transactions(
		id,
		credit_card_id,
		amount,	
		status,
		description,
		store,
		created_at) VALUES($1, $2, $3, $4, $5, $6, $7)`

	stmt, err := repo.db.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		t.ID,
		t.CreditCardId,
		t.Amount,
		t.Status,
		t.Description,
		t.Store,
		t.CreatedAt,
	)

	if err != nil {
		return err
	}

	if t.Status == "approved" {
		err = repo.updateBalance(cc)

		if err != nil {
			return err
		}
	}
	err = stmt.Close()

	if err != nil {
		return err
	}

	return nil
}

func (repo *DatabaseRepository) updateBalance(cc domains.CreditCard) error {
	_, err := repo.db.Exec(
		"UPDATE credit_cards SET balance = $1 WHERE id = $2",
		cc.Balance, cc.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repo *DatabaseRepository) CreateCreditCard(cc domains.CreditCard) error {
	query := `INSERT INTO credit_cards(
		id,
		name,
		number,
		expiration_month,
		expiration_year,
		cvv,
		balance,
		balance_limit
	) VALUES($1, $2, $3, $4, $5, $6, $7, $8)`
	stmt, err := repo.db.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		cc.ID,
		cc.Name,
		cc.Number,
		cc.ExpirationMonth,
		cc.ExpirationYear,
		cc.CVV,
		cc.Balance,
		cc.Limit,
	)

	if err != nil {
		return err
	}

	err = stmt.Close()

	if err != nil {
		return err
	}

	return nil
}

func (repo *DatabaseRepository) GetCreditCard(cc domains.CreditCard) (domains.CreditCard, error) {
	var c domains.CreditCard

	query := "SELECT id, balance, balance_limit, FROM credit_cards WHERE number=$1"
	stmt, err := repo.db.Prepare(query)

	if err != nil {
		return c, err
	}

	if err = stmt.QueryRow(cc.Number).Scan(&c.ID, &c.Balance, &c.Limit); err != nil {
		return c, errors.New("Credit Card does not exists!")
	}

	return c, nil
}
