package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func main() {
	soma := somar(10, 20)
	fmt.Println("Resultado da Função somar: ", soma)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("Texto da função 1")
}
