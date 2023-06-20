package main

import "fmt"

func main() {
	//ARITIMÉTICOS
	soma := 1 + 2
	subtracao := 1 - 2
	multiplicacao := 10 * 5
	divisao := 10 / 4
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, multiplicacao, divisao, restoDaDivisao)

	var num1 int16 = 10
	var num2 int16 = 25
	var soma2 int16 = num1 + num2
	fmt.Println(soma2)
	// FIM DOS ARITIMÉTICOS

	//ATRIBUIÇÃO
	var varriavel1 string = "String1"
	varriavel2 := "String2"
	fmt.Println(varriavel1, varriavel2)
	//FIM DOS OPERADORES ATRIBUIÇÃO

	//OPERADORES RELACIONAIS TRUE OU FALSE
	fmt.Println(1 > 2)  // false
	fmt.Println(1 >= 2) //false
	fmt.Println(1 == 2) //false
	fmt.Println(1 <= 2) //true
	fmt.Println(1 < 2)  //true
	fmt.Println(1 != 2) //true
	//FIM DOS OPERADORES RELACIONAIS

	//OPERADORES DE LÓGICOS
	fmt.Println("..................")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) // False. && analisa tudo e só retorna true se todos forem true
	fmt.Println(verdadeiro || falso) // True. || precisa de apenas um true para retorna true
	fmt.Println(!verdadeiro)         // False. ! é uma negação
	fmt.Println(!falso)              // True. ! é uma negação
	//FIM DOS OP. LÓGICOS

	//OPERADORES UNIÁRIOS
	fmt.Println("..................")
	numero := 10
	numero++     // inclementa mais 1
	numero += 15 // é mesma coisa numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 20 // é mesma coisa. numero = numero - 20
	fmt.Println(numero)

	numero *= 3 // é mesma coisa numero = numero * 3
	//FIM DE OP. UNÁRIO

}
