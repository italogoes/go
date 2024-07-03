package main

import (
	"fmt"
	"math"
)

// Interface
type forma interface {
	area() float64
}

// Structs
type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

// Implementaçoes dos structs com a interface
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

// Métodos
func escreverArea(f forma) {
	fmt.Printf("A área da forma é %f\n", f.area())
}

func main() {
	r := retangulo{10, 25}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)
}
