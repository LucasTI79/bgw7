package main

import (
	"fmt"
)

func Sum(num1, num2 int, channel chan<- int) {
	result := num1 + num2
	channel <- result
}

func main() {
	results := make(chan int)

	for i := 0; i < 100; i++ {
		go Sum(i+0, i*2, results)
	}

	// so podemos usar em canais fechados
	for value := range results {
		fmt.Println("value:", value)
	}
}
