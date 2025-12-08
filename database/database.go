package database

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type ExchangeRate struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	Bid  float64 `json:"bid,string"`
}

func NewExchangeRate(name string, price float64) *ExchangeRate {
	return &ExchangeRate{
		ID:   uuid.New().String(),
		Name: name,
		Bid:  price,
	}
}

func InsertExchangeRate(c context.Context, exchangeRate *ExchangeRate) error {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/desafio")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("insert into exchangerate(id, name, bid) values(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	ctx, cancel := context.WithTimeout(c, 10*time.Millisecond)
	if err != nil {
		panic(err)
	}
	defer cancel()

	_, err = stmt.ExecContext(ctx, exchangeRate.ID, exchangeRate.Name, exchangeRate.Bid)
	return err
}
