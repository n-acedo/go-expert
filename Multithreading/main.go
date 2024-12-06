package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"
)

type BrasilApi struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
	Fonte        string `json:"fonte"`
}

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Unidade     string `json:"unidade"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Fonte       string `json:"fonte"`
}

func main() {
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8000", nil)
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	ch1 := make(chan *BrasilApi)
	ch2 := make(chan *ViaCep)
	cep := r.URL.Query().Get("cep")

	go func() {
		adress, _ := BuscaBrasilApi(cep)
		ch1 <- adress
	}()

	go func() {
		adress, _ := BuscaViaCep(cep)
		ch2 <- adress
	}()

	select {
	case end := <-ch1:
		json.NewEncoder(os.Stdout).Encode(end)

	case end := <-ch2:
		json.NewEncoder(os.Stdout).Encode(end)

	case <-time.After(time.Second):
		println("timeout")
	}
}

func BuscaBrasilApi(cep string) (*BrasilApi, error) {
	resp, error := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()
	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}

	var address BrasilApi

	error = json.Unmarshal(body, &address)
	if error != nil {
		return nil, error
	}

	address.Fonte = "BrasilApi"

	return &address, nil
}

func BuscaViaCep(cep string) (*ViaCep, error) {
	resp, error := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()
	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}

	var address ViaCep

	error = json.Unmarshal(body, &address)
	if error != nil {
		return nil, error
	}

	address.Fonte = "ViaCep"

	return &address, nil
}
