package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "segunda-feira"
	case 3:
		return "terça-feira"
	case 4:
		return "quarta-feira"
	case 5:
		return "quinta-feira"
	case 6:
		return "sexta-feira"
	case 7:
		return "sábado"
	default:
		return "número inválido"
	}

}

func diaDaSemana2(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
	case numero == 2:
		diaDaSemana = "segunda"
	case numero == 3:
		diaDaSemana = "terça"
	case numero == 4:
		diaDaSemana = "quarta"
	case numero == 5:
		diaDaSemana = "quinta"
	case numero == 6:
		diaDaSemana = "sexta"
	case numero == 7:
		diaDaSemana = "sábado"
	default:
		diaDaSemana = "Numero invalido"

	}
	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(7)
	fmt.Println(dia)

	dia2 := diaDaSemana2(3)
	fmt.Println(dia2)

}
