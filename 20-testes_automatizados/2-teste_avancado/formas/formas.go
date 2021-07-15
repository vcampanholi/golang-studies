package formas

import (
	"math"
)

type Forma interface {
	area() float64
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

type circulo struct {
	raio float64
}

func (c circulo) Area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}
