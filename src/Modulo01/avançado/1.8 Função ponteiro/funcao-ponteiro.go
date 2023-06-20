package main

import "fmt"

// SERVE PARA FAZER UMA COPIA SEM ALTERAR A VARIAVEL
func inverteSinal(numero int) int {
	return numero * (-1)
}

// O PONTEIRO SERVE PARA ALTERA AS VARIAVEL
func inverteComPONTEIRO(numero *int) {
	*numero = *numero * -1
}

func main() {
	// AQUI ELE SÓ COPIA A VARIÁVEL , MAS NÃO ALTERA A VARAIÁVEL
	numero := 20                            // nao altera
	fmt.Println(numero)                     // 20
	numeroInvertido := inverteSinal(numero) // observe a diferença
	fmt.Println(numeroInvertido)            // -20
	fmt.Println(numero)                     // 20 .Viu ainda continua sendo 20

	fmt.Println("------------------")
	// AQUI ELE ALTERA A VARIÁVEL
	novoNumero := 40        // será alterada
	fmt.Println(novoNumero) // 40
	inverteComPONTEIRO(&novoNumero)
	fmt.Println(novoNumero) // - 40
	fmt.Println(novoNumero) // - 40
}
