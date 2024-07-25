package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup // sincronizar as goroutines/ grupo de goroutines

	waitGroup.Add(4) // aqui falo para o waitGroup que vou adicionar 2 goroutines

	go func() { // esse go na funcao é uma goroutine e faz com que a funcao seja executada em paralelo com todas as outras que tenham go na frente
		escrever("Go routine 1!") //gouroutine
		waitGroup.Done()          // aqui falo para o waitGroup que a goroutine terminou é como se fizesse waitGroup.Add(-1)
	}()

	go func() {
		escrever("Go routine 2!")
		waitGroup.Done()
	}()

	go func() { // esse go na funcao é uma goroutine e faz com que a funcao seja executada em paralelo com todas as outras que tenham go na frente
		escrever("Go routine 3!") //gouroutine
		waitGroup.Done()          // aqui falo para o waitGroup que a goroutine terminou é como se fizesse waitGroup.Add(-1)
	}()

	go func() {
		escrever("Go routine 4!")
		waitGroup.Done()
	}()

	waitGroup.Wait() // aqui falo para o waitGroup esperar todas as goroutines terminarem e chegar a zero
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
