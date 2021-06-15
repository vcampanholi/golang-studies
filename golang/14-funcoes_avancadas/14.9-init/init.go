package main

import "fmt"

var n int

//Função Init é executada sempre antes do main, pode ser uma função por arquivo
func init() {
	fmt.Println("Executando a funcao init")
	n = 10
}

func main() {
	fmt.Println("Executando a funcao main")
	fmt.Println(n)
}
