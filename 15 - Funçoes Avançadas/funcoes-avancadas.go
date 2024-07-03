package main

import "fmt"

// Funcoes com retorno nomeado
func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

// Fim

// Funcoes variatricas
func funcaoVariatrica(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func funcaoVariatrica2(texto string, numeros ...int) {
	for _, valor := range numeros {
		fmt.Println(texto, valor)
	}
}

// Fim

// Funçoes recursivas
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	soma, subtracao := calculosMatematicos(300, 100)
	fmt.Println(soma, subtracao)

	chamandoFuncaoVariatrica := funcaoVariatrica(1, 2, 3, 4, 5)
	fmt.Println(chamandoFuncaoVariatrica)

	funcaoVariatrica2("Oi, Italo", 1, 2, 3, 4, 5)

	// Funçoes anonimas
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando o valor de texto.")

	fmt.Println(retorno)

	posicao := uint(5)
	for i := uint(0); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}
