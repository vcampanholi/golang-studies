package main

import (
	"fmt"
	"time"
)

//Encapsula a chamada de uma goroutine e retorna um canal

func main() {
	//A GO rouitine fica encapsulada dentro do método escrever
	canal := escrever("Texto Canal")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
