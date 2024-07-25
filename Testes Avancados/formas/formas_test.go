package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	//TDD - Test Driven Development - Desenvolvimento Orientado a Testes - Red, Green, Refactor (Criar o teste, Fazer o teste passar, Refatorar o código)

	t.Run("Retangulo", func(t *testing.T) {
		retangulo := Retangulo{10, 12}

		areaEsperada := float64(120)
		areaRecebida := retangulo.Area()

		if areaEsperada != areaRecebida {
			t.Errorf("Area esperada %f, mas recebida %f", areaEsperada, areaRecebida)
		}
	})
	t.Run("Circulo", func(t *testing.T) {
		circulo := Circulo{10}

		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circulo.Area()

		if areaEsperada != areaRecebida {
			t.Errorf("Area esperada %f, mas recebida %f", areaEsperada, areaRecebida)
		}
	})
}
