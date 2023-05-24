package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	// Elas têm o mesmo valor. Se a gente voltar isso aqui vai ver que mudou só o valor da primeira variável.
	// Então a primeira variável teve seu valor aumentado mas a segunda a segunda não teve continua com.
	// Porque quando você atribui valor para uma variável por padrão esse valor é uma cópia.
	// Então assim aqui eu tenho o dobro 10 dessa variável e aqui eu estou jogando o valor da variável que
	// está fazendo na prática e está pegando o valor dessa variável fazendo uma cópia dele e jogando pra cá.
	// Então dessa linha pra baixo as duas variáveis não tem mais nenhuma ligação com a outra.
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println("1º amostra :", variavel1, variavel2)
	variavel2++
	fmt.Println("2º amostra :", variavel1, variavel2)

	//PONTEIRO É UMA REFERENCIA DE MEMORIA
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println("amostra do PONTEIRO :", variavel3, *ponteiro, ponteiro)

}
