package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f \n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	//MANEIRAS DIFERENTES DE FAZER A MESMA COISA

	// return (c.raio * c.raio) * 3.14   //314.16
	// return (c.raio * c.raio) * math.Pi //314.16
	return math.Pow(c.raio, 2) * math.Pi //314.16

}

func main() {
	r := retangulo{10, 15} //150.00
	escreverArea(r)

	c := circulo{10} //314.16
	escreverArea(c)

}
