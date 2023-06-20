package main

import "fmt"

// Será sempre iniciada primeiro idependedo lugar que ela estiver no cod

func init() {
	fmt.Println("Execuntando a função init")
}

func main() {
	fmt.Println("Função main sendo executada")
}
