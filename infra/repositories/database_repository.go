package repositories

import (
	"database/sql"

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
