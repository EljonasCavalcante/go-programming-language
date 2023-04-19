package main

import "fmt"

func main() {
	var variavel1 string = "variável 1"
	variavel2 := "variével 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Primeiro teste"
		variavel4 string = "Variavel teste"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, vavariavel6 := "xxxxx", "bbbbbb"

	fmt.Println(variavel5, vavariavel6)
}
