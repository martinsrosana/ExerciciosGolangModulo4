package formas

import (
	"math"
)

// Forma é uma interface que define o comportamento de uma forma geométrica
type Forma interface {
	area() float64
}

// Retangulo é uma struct que representa um retângulo
type Retangulo struct {
	Altura  float64
	Largura float64
}

// Area calcula a área do retângulo
func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

// Circulo é uma struct que representa um círculo
type Circulo struct {
	Raio float64
}

// Area calcula a área do círculo
func (c Circulo) Area() float64 {
	return math.Pi * (c.Raio * c.Raio) //ou return math.Pi * math.Pow (c.raio, 2)
}

// Quadrado é uma struct que representa um quadrado
type Quadrado struct {
	Lado float64
}

// Area calcula a área do quadrado
func (q Quadrado) Area() float64 {
	return math.Pow(q.Lado, 2)
}
