package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(4)

	go func() {
		escrever("Routine 01")
		waitGroup.Done()
	}()

	go func() {
		escrever("Routine 02")
		waitGroup.Done()
	}()

	go func() {
		escrever("Routine 03")
		waitGroup.Done()
	}()

	go func() {
		escrever("Routine 04")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever(texto string) {

	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}

}
