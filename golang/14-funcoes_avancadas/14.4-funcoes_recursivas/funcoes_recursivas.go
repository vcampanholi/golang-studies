package main

import "fmt"

//Soma das duas ultimas posições em relação a posição atual
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	posicao := uint(12)
	fmt.Println(fibonacci(posicao))

	for i := uint(1); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}

}
