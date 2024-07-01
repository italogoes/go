package main

import "fmt"

func diaDaSemana(dia int) string {
	switch dia {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	default:
		return "Dia inválido!"
	}

}

func diaDaSemana2(dia int) string {
	switch {
	case dia == 1:
		return "Domingo"
	default:
		return "Dia inválido!"
	}
}

func diaDaSemana3(dia int) string {
	var diaDaSemana string

	switch {
	case dia == 1:
		diaDaSemana = "Domingo"
	default:
		diaDaSemana = "Dia inválido!"
	}

	return diaDaSemana
}

func main() {
	dia := diaDaSemana(3)
	dia3 := diaDaSemana3(1)
	fmt.Println(dia, dia3)
}
