package main

import "fmt"

func main() {
	fmt.Println("'For com Map'")

	usuario := map[string]string{
		"nome":      "Jonas",
		"sobrenome": "Cavalcante",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
