package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	//Função Simples
	soma := somar(10, 20)
	fmt.Println(soma)

	//Função como tipo
	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	//Função com mais de um retorno
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	//Função com mais de um retorno e exibindo só um dos resultados(Usar underline no segundo resultado)
	primeiroResultado, _ := calculosMatematicos(10, 15)
	fmt.Println(primeiroResultado)
}
