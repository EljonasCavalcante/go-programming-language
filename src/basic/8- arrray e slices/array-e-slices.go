package main

import (
	"fmt"
)

func main() {
	fmt.Println("Array e Slices")
	var array1 [5]int

	array2 := [5]string{"posição0", "posição1", "posição2", "posição3", "posição4"}
	fmt.Println(array1, array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//SLICE É UMA FATIA DE UM ARRAY. APONTA PARA UMA ARRAY

	slice := []int{10, 12, 13, 14, 15}
	fmt.Println(slice) //[10 12 13 14 15]

	slice = append(slice, 18)
	fmt.Println(slice) //[10 12 13 14 15 18]

	slice2 := array2[1:3]
	fmt.Println(slice2) //[posição1 posição2]

	array2[1] = "POSICAO ALTERADA"
	fmt.Println(slice2) //[POSICAO ALTERADA posição2]

}
