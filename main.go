package main

import (
	"net/http"

	"github.com/yuri-pires/busca-cep/handlers"
)

func main() {
	// Criamos um multiplex para anexar no servidor, e nele atachamos nossas rotas.
	// Com essa abordagem, podemos ter vários servidores em uma unica aplicação
	mux := http.NewServeMux()
	// HandleFunc organiza as rotas com os handlers correspondentes
	mux.HandleFunc("/", handlers.BuscaCepHandler)
	// ListenAndServe inicia o servidor na porta 8080
	http.ListenAndServe(":8080", mux)
}
