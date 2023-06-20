package main

import "fmt"

// esse "..." pega geral
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	somaTotal := soma(1, 2, 3, 4, 6, 200, 102, 12, 13)
	fmt.Println(somaTotal)

	escrever("Olá Mundo", 1, 2, 5, 10)
}
