package main

import "fmt"

func main() {
	// Declarando de forma explicita
	var nome string = "Italo"
	// Declarando de forma implicita
	// A variavel vai se auto tipar conforme o valor que receber
	sobrenome := "Goes"
	fmt.Println(nome + " " + sobrenome)

	// Declarando varias variaveis de uma vez
	var (
		num1  int = 1
		num10 int = 10
	)
	fmt.Println("O resultado da soma é igual a:", num1+num10)

	// Atribuindo valores as variaveis separados por vírgulas
	num2, num3 := 100, 200
	fmt.Println(num2, num3)
}
