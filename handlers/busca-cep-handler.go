package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/yuri-pires/busca-cep/services"
)

// BuscaCepHandler é uma função de handler que segue o padrão (http.ResponseWriter, *http.Request)
func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Recurso inválido", http.StatusMethodNotAllowed)
		return
	}

	// Extrai o parâmetro 'cep' da URL da requisição
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Informe o CEP para realizar a busca"))
		// Early return para encerrar a função em caso de erro
		return
	}

	// Chama a função BuscarCep do pacote 'request' para buscar o CEP
	viaCepResponse := services.BuscarCep(cepParam)
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
