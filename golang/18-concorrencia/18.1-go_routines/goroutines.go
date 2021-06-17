package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Teste")
	escrever("Teste 02")
}

func escrever(texto string) {

	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}

}
