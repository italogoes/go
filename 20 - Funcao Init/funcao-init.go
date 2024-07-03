package main

import "fmt"

// A funçao init sempre vai rodar antes da main
// Ela pode ficar em varios arquivos
func init() {
	fmt.Println("Funçao init")
}

func main() {
	fmt.Println("Funçao main")
}
