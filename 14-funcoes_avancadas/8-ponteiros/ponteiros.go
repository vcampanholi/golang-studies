package main

import "fmt"

//Passando o parâmetro numero por cópia
func inverterSinal(numero int) int {
	return numero * -1
}

//Passando o parâmetro numero por referência
func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {

	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
