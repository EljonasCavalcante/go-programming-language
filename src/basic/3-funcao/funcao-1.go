package main

import "fmt"

// As funções podem ter mais de um retorno
func calaculosMatematicas(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	resultadoSoma, resultadoSubtracao := calaculosMatematicas(10, 20)
	fmt.Println("Resultado da Soma:", resultadoSoma, "- Resultado da Subtração:", resultadoSubtracao)
	// Digamos que você deseja prescise apenas um resposta, você pode usar _
	// resultadoSoma, _ := calaculosMatematicas(10, 20)
	// fmt.Println("Resultado da Soma:", resultadoSoma)
}
