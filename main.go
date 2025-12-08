package main

import (
	"fmt"

	"github.com/luanorlando/desafio-go-1/server"
)

func main() {
	fmt.Println("Hello, World!")
	server.FetchDollsExchangeRate()
}

// CREATE TABLE exchangerate (id VARCHAR(255) PRIMARY KEY, name VARCHAR(80), bid DECIMAL(10,2));
