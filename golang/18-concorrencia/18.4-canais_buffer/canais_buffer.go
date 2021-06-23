package main

import "fmt"

func main() {
	canal := make(chan string, 2)

	canal <- "Buffer 01"
	canal <- "Buffer 02"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
