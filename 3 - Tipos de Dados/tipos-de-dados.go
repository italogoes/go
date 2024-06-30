package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = 2312312312234213142
	fmt.Println(numero)

	// uint sao numeros sem sinais.
	var numero2 uint32 = 123123
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 123
	fmt.Println(numero3)

	// UINT8 = byte
	var numero4 byte = 111
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.11
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123.12
	fmt.Println(numeroReal2)

	var booleano1 bool = false
	if booleano1 == true {
		fmt.Println("Booleano é:", booleano1)
	} else {
		fmt.Println("Booleano é:", booleano1)
	}

	var erro error = errors.New("Erro interno!")
	fmt.Println(erro)
}
