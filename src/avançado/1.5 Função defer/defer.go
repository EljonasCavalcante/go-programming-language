package main

import "fmt"

func func1() {
	fmt.Println("Execuntando a função 1")
}
func func2() {
	fmt.Println("Execuntando a função 2")
}

func main() {

	//DEFER = ADIAR
	defer func1()
	func2()

}
