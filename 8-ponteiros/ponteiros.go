package main

import "fmt"

func main() {
	//Atribuindo valores por cópia
	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)
	variavel1++
	fmt.Println(variavel1, variavel2)

	//Ponteiro é uma referência de memória
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3 // & -> atribuiu o valor da variável ao endereço de memória do ponteiro

	//Exibe o endereço de memória onde foi alocado o valor da variável
	fmt.Println(variavel3, ponteiro)

	//Exibe o valor da variável no endereço da memória
	fmt.Println(variavel3, *ponteiro) //* -> Exibe o valor da variável
}
