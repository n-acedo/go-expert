package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type Quote struct {
	Usdbrl Usdbrl `json:"USDBRL"`
}

type Usdbrl struct {
	ID         string
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func main() {
	http.HandleFunc("/cotacao", DollarQuoteHandler)
	http.ListenAndServe(":8080", nil)
}

func DollarQuoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cotacao" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ctx := r.Context()

	quote, err := DollarQuote(ctx)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(quote)
}

func DollarQuote(ctx context.Context) (*string, error) {
	ctxQuote, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctxQuote, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)

	select {
	case <-ctxQuote.Done():
		log.Println("ctxQuote - Request com timeout atingido")
		return nil, ctxQuote.Err()
	default:
		log.Println("ctxQuote - Request processada dentro do tempo")
	}

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data Quote

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", "./cotacoes.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	quote := newQuote(&data)

	err = insertQuote(ctxQuote, db, quote)
	if err != nil {
		if err.Error() == "context deadline exceeded" {
			log.Println("ctxInsert - Timeout atingido na persistência dos dados no banco")
		}
		return nil, err
	}

	return &data.Usdbrl.Bid, nil
}

func newQuote(quote *Quote) *Usdbrl {
	usdbrl := quote.Usdbrl

	return &Usdbrl{
		ID:         uuid.New().String(),
		Code:       usdbrl.Code,
		Codein:     usdbrl.Codein,
		Name:       usdbrl.Name,
		High:       usdbrl.High,
		Low:        usdbrl.Low,
		VarBid:     usdbrl.VarBid,
		PctChange:  usdbrl.PctChange,
		Bid:        usdbrl.Bid,
		Ask:        usdbrl.Ask,
		Timestamp:  usdbrl.Timestamp,
		CreateDate: usdbrl.CreateDate,
	}
}

func insertQuote(ctx context.Context, db *sql.DB, quote *Usdbrl) error {
	ctxInsert, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	stmt, err := db.Prepare("insert into quotes(id, code, code_in, name, high, low, var_bid, pct_change, bid, ask, timestamp, create_date) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctxInsert, quote.ID, quote.Code, quote.Codein, quote.Name, quote.High, quote.Low, quote.VarBid, quote.PctChange, quote.Bid, quote.Ask, quote.Timestamp, quote.CreateDate)
	if err != nil {
		return err
	}

	return nil
}
