package main

import "fmt"

// operations performs basic arithmetic operations on two float64 values
func operations(value1, value2 float64) (float64, float64, float64, float64) {
	summatory := value1 + value2
	subtraction := value1 - value2
	multiplication := value1 * value2

	var division float64
	if value2 != 0 {
		division = value1 / value2
	}

	return summatory, subtraction, multiplication, division
}

func main() {
	summatory, subtraction, multiplication, division := operations(10.0, 5.0)

	fmt.Println("summatory:", summatory)
	fmt.Println("subtraction:", subtraction)
	fmt.Println("multiplication:", multiplication)
	fmt.Println("division:", division)
}
