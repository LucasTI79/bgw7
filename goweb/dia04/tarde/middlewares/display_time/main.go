package main

import (
	"fmt"
	"time"
)

// assinatura da funcao
type operation = func(int, int) int

func displayTime(op operation) operation {
	return func(num1, num2 int) int {
		fmt.Println("time:", time.Now())
		return op(num1, num2)
	}
}

func sum(num1, num2 int) int {
	return num1 + num2
}

func sub(num1, num2 int) int {
	return num1 - num2
}

func main() {
	sumWithDisplayTime := displayTime(sum)
	fmt.Println(sumWithDisplayTime(1, 2))

	subWithDisplayTime := displayTime(sub)
	fmt.Println(subWithDisplayTime(1, 2))
}
