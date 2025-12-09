package server

import (
	"context"
	"encoding/json"
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
		http.Error(w, "Falha ao preparar Request", http.StatusInternalServerError)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Erro ao codificar JSON", http.StatusNotFound)
	}
	defer res.Body.Close()

	var resp apiResponse
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		http.Error(w, "Erro ao Decodificar JSON response", http.StatusInternalServerError)
	}

	exchangeRate := resp.USDBRL
	err = database.InsertExchangeRate(ctx, database.NewExchangeRate(exchangeRate.Name, exchangeRate.Bid))
	if err != nil {
		http.Error(w, "Erro ao salvar no banco de dados", http.StatusInternalServerError)
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]float64{"cotacao": exchangeRate.Bid})
	if err != nil {
		http.Error(w, "Erro ao codificar JSON", http.StatusInternalServerError)
	}
}
