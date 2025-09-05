package main

import (
	"fmt"

	"github.com/lucasti79/bgw7/tdd/factorial"
)

func main() {
	arg := 3
	result := factorial.Factorial(arg)

	fmt.Println(result)
}
