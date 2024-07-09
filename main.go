package main

import (
	"encoding/json"
	"net/http"

	"github.com/yuri-pires/busca-cep/request"
)

func main() {
	// HandleFunc organiza as rotas com os handlers correspondentes
	http.HandleFunc("/", BuscaCepHandler)
	// ListenAndServe inicia o servidor na porta 8080
	http.ListenAndServe(":8080", nil)
}

// BuscaCepHandler é uma função de handler que segue o padrão (http.ResponseWriter, *http.Request)
func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	// Extrai o parâmetro 'cep' da URL da requisição
	cepParam := r.URL.Query().Get("cep")

	// Verifica se o parâmetro 'cep' está presente na requisição
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Informe o CEP para realizar a busca"))
		// Early return para encerrar a função em caso de erro
		return
	}

	// Chama a função BuscarCep do pacote 'request' para buscar o CEP
	viaCepResponse := request.BuscarCep(cepParam)
	// Converte a resposta da função BuscarCep (uma struct) para JSON
	cepJson, err := json.Marshal(viaCepResponse)

	// Verifica se houve erro na conversão para JSON
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao consultar o CEP para realizar a busca"))
		// Early return para encerrar a função em caso de erro
		return
	}

	// Define o header 'Content-Type' como 'application/json'
	w.Header().Set("Content-Type", "application/json")
	// Define o status HTTP como 200 OK
	w.WriteHeader(http.StatusOK)
	// Escreve a resposta JSON no corpo da resposta HTTP
	w.Write(cepJson)
	// Alternativamente, poderíamos usar o encoder JSON para escrever a resposta:
	// json.NewEncoder(w).Encode(viaCepResponse)
}
