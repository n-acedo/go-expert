package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)

	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if ctx.Err() != nil {
		log.Println("ctx - Timeout atingido na solicitacao da cotacao")
		panic(err)
	}

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	file, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	sentence := fmt.Sprintf("Dólar: {%s}", body)
	_, err = file.Write([]byte(sentence))
	if err != nil {
		panic(err)
	}
}
