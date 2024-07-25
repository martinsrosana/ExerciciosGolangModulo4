package main

import (
	"fmt"
	"time"
)

func main() {

	canal := escrever("Olá Mundo!") // aqui a gente recebe o canal que a funcao escrever retornou
	for i := 0; i < 10; i++ {       // aqui a gente faz um loop para receber os dados do canal
		fmt.Println(<-canal)
	}
}

// a funcao escrever esta cuidando de driar o canal e de enviar os dados para o canal
func escrever(texto string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s ", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
