package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando I")
	// 	time.Sleep(time.Second)

	// }
	// fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando J", j)
	}

	// Range
	nomes := [3]string{"Joao", "Davi", "Lucas"}

	for indice, valor := range nomes {
		fmt.Println(indice, valor)
	}

	for indice, valor := range "PALAVRA" {
		fmt.Println(indice, string(valor))
	}

	// FOR em maps
	usuario := map[string]string{
		"nome":      "Italo",
		"sobrenome": "Silva",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// For infinito
	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}

}
