package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Print("Fizz")
			fallthrough
		case i%5 == 0:
			fmt.Print("Buzz")
		case i%3 == 0:
			fmt.Print("Fizz")
		default:
			fmt.Printf("%d", i)
		}
		fmt.Print(" ")
	}
}

// case "x":
//     fallthrough
// case "y":
//     funcaoTal()
