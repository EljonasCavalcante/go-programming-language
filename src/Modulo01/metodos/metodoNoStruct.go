package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("-- Salvando os dados do Usuário %s no banco de dados -- ", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}
func main() {
	usuario1 := usuario{"Usuário", 20}
	usuario1.salvar()

	usuario2 := usuario{"Usuário", 30}
	usuario2.salvar()

	maiorDeIdadde := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdadde)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
