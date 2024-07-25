package main

// canal pq é um canal de comunicação e usamos para sincronizar as gourotines

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)       // aqui eu crio um canal de string
	go escrever("Olá Mundo!", canal) //gouroutine

	fmt.Println("Depois da função escrever ser chamada")

	for {
		mensagem, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}
	fmt.Println("Fim do programa")
}

// deadlock é qdo vc nao tem mais nenhum dado sendo enviado para o seu canal e vc está esperando por um dado que nunca vai chegar e o programa fica travado
// para resolver o deadlock vc pode usar o close(canal) para fechar o canal
// o close(canal) só pode ser usado no canal de envio
// o close(canal) só pode ser usado uma vez
// o close(canal) só pode ser usado pelo emissor do canal
// o close(canal) só pode ser usado no canal que foi criado com make
// o close(canal) só pode ser usado no canal que não é um canal bidirecional
// o close(canal) só pode ser usado no canal que não é um canal de recebimento
// o close(canal) só pode ser usado no canal que não é um canal de envio
// o close(canal) só pode ser usado no canal que não é um canal de leitura
// o close(canal) só pode ser usado no canal que não é um canal de escrita
// o close(canal) só pode ser usado no canal que não é um canal de leitura e escrita
// o close(canal) só pode ser usado no canal que não é um canal de escrita e leitura
// o close(canal) só pode ser usado no canal que não é um canal de leitura e envio
// o close(canal) só pode ser usado no canal que não é um canal de envio e leitura
// o close(canal) só pode ser usado no canal que não é um canal de escrita e recebimento
// o close(canal) só pode ser usado no canal que não é um canal de recebimento e escrita
// o close(canal) só pode ser usado no canal que não é um canal de leitura e recebimento
// o close(canal) só pode ser usado no canal que não é um canal de recebimento e leitura
// o close(canal) só pode ser usado no canal que não é um canal de escrita e leitura e recebimento
// o close(canal) só pode ser usado no canal que não é um canal de escrita e recebimento e leitura
// o close(canal) só pode ser usado no canal que não é um canal de leitura e escrita e recebimento
// o close(canal) só pode ser usado no canal que não é um canal de leitura e recebimento e escrita
// o close(canal) só pode ser usado no canal que não é um canal de recebimento e escrita e leitura
// o close(canal) só pode ser usado no canal que não é um canal de recebimento e leitura e escrita
// o close(canal) só pode ser usado no canal que não é um canal de escrita e leitura e envio
// o close(canal) só pode ser usado no canal que não é um canal de escrita e envio e leitura
// o close(canal) só pode ser usado no canal que não é um canal de leitura e escrita e envio
// o close(canal) só pode ser usado no canal que não é um canal de leitura e envio e escrita
// o close(canal) só pode ser usado no canal que não é um canal de envio e leitura e escrita
// o close(canal) só pode ser usado no canal que não é um canal de envio e escrita e leitura
// o close(canal) só pode ser usado no canal que não é um canal de escrita e leitura e recebimento e envio
// o close(canal) só pode ser usado no canal que não é um canal de escrita e recebimento e leitura e envio
// o close(canal) só pode ser usado no canal que não é um canal de leitura e escrita e recebimento e envio
// o close(canal) só pode ser usado no canal que não é um canal de leitura e recebimento e escrita e envio
// o close(canal) só pode ser usado no canal que não é um canal de recebimento e escrita e leitura e envio
// o close(canal) só pode ser usado no canal que não é um canal de recebimento e leitura e escrita e envio
// o close(canal) só pode ser usado no canal que não é um canal de escrita e leitura e envio e recebimento
// o close(canal) só pode ser usado no canal que não é um canal de escrita e envio e leitura e recebimento
// o close(canal) só pode ser usado no canal que não é um canal de leitura e escrita e envio e recebimento
// o close(canal) só pode ser usado no canal que não é um canal de leitura e envio e escrita e recebimento
// o close(canal) só pode ser usado no canal que não é um canal de envio e leitura e escrita e recebimento
// o close(canal) só pode ser usado no canal que não é um canal de envio e escrita e leitura e recebimento
// o close(canal) só pode ser usado no canal que não é um canal de escrita e leitura e recebimento e envio
// o close(canal) só pode ser usado no canal que não é um canal de escrita e recebimento e leitura e envio
// o close(canal) só pode ser usado no canal que não é um canal de leitura e escrita e recebimento e envio
// o close(canal) só pode ser usado no canal que não é um canal de leitura e envio e recebimento e escrita
// o close(canal) só pode ser usado no canal que não é um canal de envio e leitura e recebimento e escrita
// o close(canal) só pode ser usado no canal que não é um canal de envio e escrita e leitura e recebimento
// o close(canal) só pode ser usado no canal que não é um canal de escrita e leitura e envio e recebimento
// o close(canal) só pode ser usado no canal que não é um canal de escrita e envio e leitura e recebimento
// o close(canal) só pode ser usado no canal que não é um canal de leitura e escrita e envio e recebimento
// o close(canal) só pode ser usado no canal que não é um canal de leitura e envio e escrita e recebimento

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}
