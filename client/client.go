package client

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

func RunClient() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Println("DEBUG: Request erro")
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("DEBUG: erro ao pegar response")
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("DEBUG: erro ao pegar body do response")
		panic(err)
	}

	log.Println("DEBUG: erro ao pegar body do response", string(body))

	saveResponseBody(body)
}

func saveResponseBody(b []byte) {

}
