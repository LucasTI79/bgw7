package main

import "fmt"

func addition(values ...float64) float64 {
	var result float64
	for _, value := range values {
		result += value
	}
	return result
}

func main() {
	result := addition(1.0, 2.0, 3.0, 4.0, 5.0)
	fmt.Println("The result of the addition is:", result)
}
