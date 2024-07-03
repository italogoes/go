package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int
}

func main() {
	var u usuario

	u.nome = "Italo"
	u.idade = 24
	fmt.Println(u)

	enderecoDoItalo := endereco{"Rua arembepe", 301}

	usuario2 := usuario{"Joao", 21, enderecoDoItalo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Davi"}
	fmt.Println(usuario3)
}
