package main

import "fmt"

func main() {
	// IF ELSE
	// IF init
	numero := -2

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Sim")
	} else if outroNumero < -1{
		fmt.Println("Nao")
	}
}
