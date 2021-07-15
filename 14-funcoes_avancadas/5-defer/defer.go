package main

import "fmt"

func funcao1() {
	fmt.Println("Função 1")
}

func funcao2() {
	fmt.Println("Função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	//Vai ser executado antes de qualuer um dos dois retornos
	defer fmt.Println("Média calculada. Resultado será retornado")
	fmt.Println("Entrando na função para verificr se o aluno está aprovado")

	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false

}

func main() {
	//DEFER adia a execução da função
	defer funcao1()
	funcao2()

	fmt.Println(alunoEstaAprovado(7, 8))
}
