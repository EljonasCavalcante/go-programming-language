package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("-- Salvando usuario %s \n", u.nome)
}

func (u usuario) maiordeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {

	usuario1 := usuario{"Manuel", 10}
	usuario1.salvar()

	eMaiordeIdade := usuario1.maiordeIdade()
	if eMaiordeIdade {
		fmt.Println("Usuario tem  maior")
	} else {
		fmt.Println("Usuario tem  menor")
	}

}
