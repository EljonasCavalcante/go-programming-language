package main

// O que é Estrutura?
// Na programação Go, uma estrutura ou struct é um tipo definido pelo usuário para armazenar uma coleção de diferentes campos em um único campo.https://www.simplilearn.com/tutorials/golang-tutorial/golang-struct

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}
type endereco struct {
	logradoro string
	numero    uint8
}

func main() {
	fmt.Println("Arquivo Structs")
	var u usuario
	u.nome = "Jonas"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua do imperio", 112}

	u2 := usuario{"Jonas", 21, enderecoExemplo}
	fmt.Println("Teste 01", u2)

	u3 := usuario{nome: "Davi"} // repare que eu chamei somente um item assim.
	fmt.Println("Teste 02", u3)

}
