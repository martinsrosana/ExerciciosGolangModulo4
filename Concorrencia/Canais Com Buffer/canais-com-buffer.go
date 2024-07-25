package main

import "fmt"

// a gente usa o buffer para evitar o deadlock mas o buffer tem um limite e se o limite for atingido o programa vai travar
// o buffer é uma fila que armazena os dados que são enviados para o canal

func main() {
	canal := make(chan string, 2) //1 esse (2) é o buffer do canal
	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"
	//canal <- "Programando em Go!" // se manter essa linha dá deadlock

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
