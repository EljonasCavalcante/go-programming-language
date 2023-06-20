// Uma função recursiva em Go é uma função que chama a si mesma. As funções recursivas podem ser muito úteis para resolver problemas que podem ser divididos em subproblemas menores e semelhantes, como calcular o fatorial de um número, gerar a sequência de Fibonacci ou percorrer uma estrutura de dados em árvore.
// Funções anônimas na Golang são funções que não estão vinculadas a um nome específico. Elas são úteis em várias situações, incluindo: Como uma função de curta duração que só é necessária em um lugar.
package main

import "fmt"

// A sucessão de Fibonacci é uma sequência de números inteiros iniciados por zero e um, no qual cada termo subsequente corresponde a soma dos dois números anteriores: 0,1, 1, 2, 3, 5, 8, 13, 21,

// é importante dizer para a funcao deve parar.
// fib (n – 1) + fib (n – 2)

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-1) + fibonacci(posicao-2)
}

func main() {
	fmt.Println("Função Recursivas")

	posicao := uint(5)

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}
	fmt.Println("Na posição 6 está o nº:", posicao)

}
