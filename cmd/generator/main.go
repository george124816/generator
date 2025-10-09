package main

import (
	"fmt"
	"os"

	"github.com/george124816/generator/internal/documents"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`should be an argument:
		- cpf
		- cnpj
		`)
	} else {
		switch os.Args[1] {
		case "cpf":
			fmt.Println(documents.GenerateCpf())
		case "cnpj":
			fmt.Println(documents.GeneraterCnpj())
		}
	}

}
