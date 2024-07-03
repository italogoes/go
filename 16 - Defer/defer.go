package main

import "fmt"

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada.")
	fmt.Println("Verificando se o aluno está aprovado...")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	fmt.Println(alunoAprovado(10, 9))
}
