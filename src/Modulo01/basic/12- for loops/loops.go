package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	//PRIMEIRO EXEMPLOS
	i := 0
	for i < 5 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second) // dormir por um segundo

	}
	fmt.Println(i)

	// SEGUNDO EXEMPLO

	for j := 0; j < 5; j++ {
		fmt.Println("Incrementando J", j)
		time.Sleep(time.Second) // dormir por um segundo
	}

	// TERCEIRO EXEMPLO

	nomes := [3]string{"João", "Jonas", "Maria"}
	for indice, nome := range nomes { // range percorre
		fmt.Println(indice, nome) // 0 João 1 Jonas 2 Maria
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))

	}

}
