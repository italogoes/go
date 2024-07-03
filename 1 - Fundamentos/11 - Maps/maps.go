package main

import "fmt"

func main() {
	// Maps

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario["nome"])

	teste := map[int]string{
		1: "Posiçao 1",
		2: "Posiçao 2",
	}
	fmt.Println(teste[1])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Joao",
			"segundo":  "Pedro",
		},
		"curso": {
			"nome": "engenharia",
		},
	}
	fmt.Println("Nome do usuário:", usuario2["nome"]["primeiro"])

	usuario2["avaliacao"] = map[string]string {
		"nome": "Ótimo Aluno",
	}
	fmt.Println("Avaliçao do estudante:", usuario2["avaliacao"]["nome"])

	delete(usuario2, "nome")
	fmt.Println(usuario2)
}
