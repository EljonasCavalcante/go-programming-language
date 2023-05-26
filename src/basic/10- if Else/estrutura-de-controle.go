package main

import "fmt"

func main() {
	fmt.Println("--Estrutura de controle--")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número maior que zero")
	} else {
		fmt.Println("Número menor que zero")
	}
}
