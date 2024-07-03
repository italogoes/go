package main

import "fmt"

func main() {
	// Arrays internos
	// A capacidade máxima é opcional

	slice := make([]float32, 10, 11)
	fmt.Println(slice)
	fmt.Println("Tamanho do slice:", len(slice)) // Tamanho
	fmt.Println("Capaciade máxima:", cap(slice)) // Capaciade maxima

	slice = append(slice, 100, 123)
	fmt.Println(slice)

	fmt.Println("Tamanho do slice:", len(slice)) // Tamanho
	fmt.Println("Capaciade máxima:", cap(slice)) // Capaciade maxima
}
