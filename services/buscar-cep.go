package services

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
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
	var viaCepStruct ViaCepResponse
	viaCepUrl := "https://viacep.com.br/ws/" + cep + "/json/"

	res, err := http.Get(viaCepUrl)
	if err != nil {
		log.Fatalf("Erro ao realizar consulta no WS Via Cep %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Erro ao ler a resposta: %v", err)
	}

	err = json.Unmarshal(body, &viaCepStruct)

	if err != nil {
		log.Fatalf("Erro ao transformar resposta: %v", err)
	}

	return &viaCepStruct
}
