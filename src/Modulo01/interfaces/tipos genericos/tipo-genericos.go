package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)

}

func main() {
	generica("string")
	generica(1)
	generica(true)
	generica(1 > 1)
	generica(1 >= 1)
	generica(1 <= 1)
	generica(nil)
}
