package main

import "fmt"

// nao podemos criar variaveis comd declaracao fora do escopo de funçõesß
// i := 0

func main() {
	// declarando variavel i
	// var i int

	// atribuindo valor a variavel i
	// i = 0

	// declarando e inicializando a variavel i
	// var i int = 0

	// declaracao curta de variavel
	// i := 0

	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}
