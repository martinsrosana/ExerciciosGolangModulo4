package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	cenarioDeTeste := []cenarioDeTeste{
		{"Rua dos Imigrantes", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Estrada do M'Boi Mirim", "Estrada"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Tipo Inválido"},
		{"RUA DOS IMIGRANTES", "Rua"},
		{"", "Tipo Inválido"},
		{"AVENIDA SAO PAULO", "Avenida"},
	}

	for _, cenario := range cenarioDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", cenario.retornoEsperado, retornoRecebido)
		}
	}
}
