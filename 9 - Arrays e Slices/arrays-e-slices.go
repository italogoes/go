package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Array
	var array1 [5]int
	array1[0] = 100
	fmt.Println(array1)

	array2 := [5]string{"Posiçao 1", "Posiçao 2", "Teste", "Teste 2", "Teste 3"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// Slice
	slice := []int{10, 20, 30}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	// Adicionando um novo valor ao slice
	slice = append(slice, 18)
	fmt.Println(slice)

	// Fatiando o array e salvando no slice
	slice2 := array2[1:4]
	fmt.Println(slice2)

	// Alterando o valor do array
	array2[1] = "Posiçao alterada"
	fmt.Println(array2)
}
