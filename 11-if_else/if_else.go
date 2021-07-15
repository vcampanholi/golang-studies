package main

import "fmt"

func main() {

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	//Estrutyura com IF INIT -> Cria a variável e atribui um valor(outroNumero := numero) e avalia a condição(outroNumero > 0)
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que 0")
	} else if numero < -10 {
		fmt.Println("Número é menor que -10")
	} else {
		fmt.Println("Está entre 10 e -10")
	}

}
