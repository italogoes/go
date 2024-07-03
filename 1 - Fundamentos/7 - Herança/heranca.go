package main

import "fmt"

type pessoa struct {
	nome   string
	idade  uint8
	altura uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	pessoa1 := pessoa{"Italo", 24, 180}
	fmt.Println(pessoa1)

	estudante1 := estudante{pessoa1, "Sistemas para internet", "Wyden"}
	fmt.Println(estudante1)
}
