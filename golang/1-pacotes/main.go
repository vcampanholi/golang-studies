package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"modulo/auxiliar"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")

	auxiliar.Escrever()
	checkmail.ValidateFormat("vando.zc@gmail.com")
}
