package main

import "fmt"

// Caso a funçao tenha retorno, precisa esfecificar o tipo de retorno
// após os parenteses
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func usuario(nome, sobrenome, email string) (string, string, string) {
	userNome := nome
	userSobrenome := sobrenome
	userEmail := email

	return userNome, userSobrenome, userEmail
}

func main() {
	// Funcao cadastro de usuário.
	nome, sobrenome, email := usuario("Italo", "Goes", "italo@gmail.com")
	fmt.Println(nome, sobrenome, email)

	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(texto string) string {
		fmt.Println(texto)
		return texto
	}

	resultado := f("Meu texto!")
	fmt.Println(resultado)

	resultadoSoma, _ := calculosMatematicos(10, 10)
	fmt.Println(resultadoSoma)
}
