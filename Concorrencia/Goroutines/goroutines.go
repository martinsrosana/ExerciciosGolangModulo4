package main

import (
	"fmt"
	"time"
)

func main() {

	escrever("Olá Mundo!")         //gouroutine
	escrever("Programando em Go!") // esse nunca vai entrar pq o de cima nunca termina
}

// assim não entra no segundo escrever
func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
