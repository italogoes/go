package main

import "fmt"

type usuario struct {
	nome  string
	idade int8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados.\n", u.nome)
}

func (u usuario) deletar() {
	fmt.Printf("Usuário %s deletado com sucesso!", u.nome)
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Italo", 24}
	usuario1.salvar()
	usuario1.deletar()
	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)
}
