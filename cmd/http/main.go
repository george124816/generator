package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/george124816/generator/internal/documents"
)

const port = ":4000"

func main() {
	handler := http.NewServeMux()

	handler.HandleFunc("/cpf", handleCpf)
	handler.HandleFunc("/cnpj", handleCnpj)

	log.Printf("serving at port %s\n", port)
	http.ListenAndServe(port, handler)
}

func handleCpf(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, documents.GenerateCpf())
}

func handleCnpj(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, documents.GeneraterCnpj())
}
