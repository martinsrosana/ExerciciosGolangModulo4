package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(escrever("Olá Mundo!"), escrever("Programando em Go!"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

// funcao multiplexar vai receber um canal de canais de strings com dois canais de entrada e um de saida
func multiplexar(canalEntrada1, canalEntrada2 <-chan string) <-chan string {
	canalSaida := make(chan string)
	go func() {
		for {
			select {
			case mensagem := <-canalEntrada1:
				canalSaida <- mensagem
			case mensagem := <-canalEntrada2:
				canalSaida <- mensagem
			}
		}
	}()

	return canalSaida

}

// a funcao escrever esta cuidando de criar o canal e de enviar os dados para o canal
// essa funcao vai retornar um canal que vai receber strings
// o canal vai ser um canal somente de leitura
func escrever(texto string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s ", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
