package main

import "fmt"

func main() {
	//FUNÇÃO ANÔNIMA
	// func(texto string) {
	// 	fmt.Println(texto)
	// }("Passando parametro")

	retornoDafuncaoAnonima := func(texto string) string {
		// return fmt.Sprintf("Recebido -> %s %d", texto, 10)
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando parametro")

	fmt.Println(retornoDafuncaoAnonima) //Recebido -> Passando parametro
}
