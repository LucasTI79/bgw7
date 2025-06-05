package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("./myFile.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	// saida padrao de sucesso do programa
	os.Exit(0)
}
