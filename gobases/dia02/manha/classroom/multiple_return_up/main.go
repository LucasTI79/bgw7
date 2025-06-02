package main

import "fmt"

func operations(value1, value2 float64) (sum, sub, mult, div float64) {
	sum = value1 + value2
	sub = value1 - value2
	mult = value1 * value2

	if value2 != 0 {
		div = value1 / value2
	}

	return
}

func main() {
	summatory, subtraction, multiplication, division := operations(10.0, 5.0)

	fmt.Println("summatory:", summatory)
	fmt.Println("subtraction:", subtraction)
	fmt.Println("multiplication:", multiplication)
	fmt.Println("division:", division)
}
