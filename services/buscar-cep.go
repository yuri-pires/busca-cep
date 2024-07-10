package services

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type ViaCepResponse struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

func BuscarCep(cep string) *ViaCepResponse {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	viaCepUrl := "https://viacep.com.br/ws/" + cep + "/json/"
	// Ao iniciar uma nova request com erro, teremos ou um Ponteiro
	// com uma requisição ou um erro.
	// Está requisição seguirá nosso contexto com timeout de 1 minuto.
	req, err := http.NewRequestWithContext(ctx, "GET", viaCepUrl, nil)
	if err != nil {
		panic(err)
	}

	// Como criamos somente a request acima, precisamos executar ela agora
	// Utilizaremos o método Do para obter uma Response ou um erro.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Erro ao realizar consulta no WS Via Cep %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erro ao ler a resposta: %v", err)
	}

	var viaCepStruct ViaCepResponse
	if err := json.Unmarshal(body, &viaCepStruct); err != nil {
		log.Fatalf("Erro ao transformar resposta: %v", err)
	}

	return &viaCepStruct
}
