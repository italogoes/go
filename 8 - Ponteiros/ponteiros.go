package main

import "fmt"

func main() {
	// O problema que é gerado sem o uso de ponteiros.
	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	// Resolvendo esse problema com ponteiros.
	// PONTEIRO É UMA REFERENCIA DE MEMORIA.
	var valorRecebido int
	var memoriaAlocada *int

	valorRecebido = 200
	memoriaAlocada = &valorRecebido
	fmt.Println(valorRecebido, memoriaAlocada)

	valorRecebido = 300
	fmt.Println(valorRecebido, memoriaAlocada)
}
