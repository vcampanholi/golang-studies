package main

import (
	"fmt"
	"time"
)

func iteracaoSimples() {
	i := 0
	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}
	fmt.Println(i)
}

func iteracaoComVariaveisInicializadas() {
	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando J", j)
		time.Sleep(time.Second)
	}
}

func iteracaoEmArray() {
	nomes := [3]string{"João", "Davi", "Lucas"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	//Sem informar o indice do FOR
	for _, nome := range nomes {
		fmt.Println(nome)
	}

}

func iteracaoEmString() {
	for indice, letra := range "PALAVRA" {
		//Retorna os valores da tabela ASC
		fmt.Println(indice, letra)
		//Converte para o caraccter da tabela ASC
		fmt.Println(indice, string(letra))
	}
}

func iteracaoEmStructs() {
	usuario := map[string]string{
		"nome":      "João",
		"sobrenome": "Paulo",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}

func main() {
	iteracaoSimples()
	iteracaoComVariaveisInicializadas()
	iteracaoEmArray()
	iteracaoEmString()
	iteracaoEmStructs()
}
