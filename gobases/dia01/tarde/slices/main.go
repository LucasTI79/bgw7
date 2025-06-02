package main

import "fmt"

func main() {
	// declarando slice
	// var a []int

	// a := []int{1, 2, 3}

	// declarando slice com declaracao curta
	a := make([]int, 5)

	// adicionando itens ao slice
	a = append(a, 1, 2, 3)

	// printando valores do slice
	fmt.Println(a)

	// obtendo capacidade
	fmt.Println(cap(a))

	// obtendo tamanho
	fmt.Println(len(a))
}
