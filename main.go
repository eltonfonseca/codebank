package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/eltonfonseca/codebank/domains"
	"github.com/eltonfonseca/codebank/infra/repositories"
	"github.com/eltonfonseca/codebank/usecase"
	_ "github.com/lib/pq"
)

func main() {
	db := setupDB()
	defer db.Close()

	cc := domains.NewCreditCard()
	cc.Number = "1234"
	cc.Name = "Elton"
	cc.ExpirationYear = 2021
	cc.ExpirationMonth = 07
	cc.CVV = 123
	cc.Limit = 1000
	cc.Balance = 0

	repo := repositories.NewDatabaseRepository(db)
	err := repo.CreateCreditCard(*cc)

	if err != nil {
		fmt.Println(err)
	}
}

func setupUseCase(db *sql.DB) usecase.UseCaseTransaction {
	repo := repositories.NewDatabaseRepository(db)
	useCase := usecase.NewUseCaseTransaction(repo)

	return useCase
}

func setupDB() *sql.DB {
	config := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"db", "5432", "postgres", "root", "codebank",
	)
	db, err := sql.Open("postgres", config)

	if err != nil {
		log.Fatal("Error on connect to database!")
	}

	return db
}
