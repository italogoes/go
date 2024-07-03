package main

import "fmt"

func recuperarFuncao() {
	if r := recover(); r != nil {
		fmt.Println("Execucao recuperada com sucesso!!!")
	}
}

func alunoTeste(n1, n2 float32) bool {
	defer recuperarFuncao()
	media := (n1 + n2) / 2

	if media < 6 {
		return false
	} else if media > 6 {
		return true
	}

	panic("A média é exatamente 6!")
}

func main() {
	fmt.Println(alunoTeste(6, 6))
}
