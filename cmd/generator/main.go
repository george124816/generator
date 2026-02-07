package main

import (
	"fmt"
	"os"

	utils "github.com/george124816/generator/internal"
	"github.com/george124816/generator/internal/documents"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`should be an argument:
		- cpf
		- cnpj
		- uuid
		`)
	} else {
		switch os.Args[1] {
		case "cpf":
			fmt.Println(documents.GenerateCpf())
		case "cnpj":
			fmt.Println(documents.GeneraterCnpj())
		case "uuid":
			fmt.Println(utils.GenerateUuid())
		}
	}

}
