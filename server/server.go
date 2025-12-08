package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/luanorlando/desafio-go-1/database"
)

type apiResponse struct {
	USDBRL database.ExchangeRate `json:"USDBRL"`
}

func FetchDollsExchangeRate() {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	var url string = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var resp apiResponse
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		panic(err)
	}

	exchangeRate := resp.USDBRL
	err = database.InsertExchangeRate(database.NewExchangeRate(exchangeRate.Name, exchangeRate.Bid))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
