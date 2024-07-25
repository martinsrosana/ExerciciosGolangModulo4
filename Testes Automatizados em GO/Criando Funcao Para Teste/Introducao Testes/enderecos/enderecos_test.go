package enderecos

// poderia ser assim também package enderecos_test
import "testing"

// poderia ser assim também import "introducao-testes/enderecos" ou importar com alias ."introducao-testes/enderecos" que ai não precisa usar o pacote antes da funcao

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	//t.Parallel() - para rodar em paralelo ou seja todos os testes vão rodar ao mesmo tempo

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
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido) //aqui se tivesse usando o pacote enderecos poderia ser enderecos.TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", cenario.retornoEsperado, retornoRecebido)
		}
	}
}

func TestQualquer(t *testing.T) {
	//	t.Parallel() - para rodar em paralelo ou seja todos os testes vão rodar ao mesmo tempo
	if 1 > 2 {
		t.Errorf("Erro!")
	}
}

// para rodar todos os testes a gente usa o comando go test ./...
// para rodar o teste a gente usa o comando go test
// para rodar o teste em paralelo a gente usa o comando go test -p 3
// para rodar o teste e ver o coverage a gente usa o comando go test --cover
// para rodar o teste e ver o coverage em html a gente usa o comando go test --coverprofile cobertura.html
// para rodar o teste e ver o coverage em html a gente usa o comando go test --coverprofile cobertura.html && go tool cover --html=cobertura.html
