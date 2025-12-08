package server

import (
	"bytes"
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

func RunServer() {
	http.HandleFunc("/cotacao", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
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
	jsonStr := fmt.Sprintf(`{ "cotacao": "%v"}`, exchangeRate.Bid)
	jsonVar := bytes.NewBuffer([]byte(jsonStr))
	w.Write(jsonVar.Bytes())

	err = database.InsertExchangeRate(ctx, database.NewExchangeRate(exchangeRate.Name, exchangeRate.Bid))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
