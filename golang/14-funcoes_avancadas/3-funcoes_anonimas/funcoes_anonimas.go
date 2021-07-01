package main

import "fmt"

func main() {

	//Funçao sem parâmetro
	func() {
		fmt.Println("Função Anônima")
	}()

	//Funçao com passagem de parâmetro e retorno
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando parâmetro")

	fmt.Println(retorno)

}
