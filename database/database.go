package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type ExchangeRate struct {
	ID   string
	Name string
	Bid  float32
}

func NewProduct(name string, price float32) *ExchangeRate {
	return &ExchangeRate{
		ID:   uuid.New().String(),
		Name: name,
		Bid:  price,
	}
}

func InsertExchangeRate(db *sql.DB, exchangeRate *ExchangeRate) error {
	db, err := sql.Open("exchangeRate", "root:root@tcp(localhost:3307)/desafio1")
	if err != nil {
		return err
	}

	defer db.Close()

	stmt, err := db.Prepare("insert into exchangeRate(id, name, bid) values(?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(exchangeRate.ID, exchangeRate.Name, exchangeRate.Bid)
	return err

}
