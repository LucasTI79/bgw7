package main

import "fmt"

func main() {
	var age uint8 = 18

	switch {
	case age >= 150:
		fmt.Println("Are you immortal?")
	case age >= 18:
		fmt.Println("You are an adult!")
	default:
		fmt.Println("You’re not an adult yet")
	}
}
