package main

import "fmt"

func main() {
	animals := []string{
		"cow",
		"dog",
		"hawk",
	}
	fmt.Println("only flies on: ", animals[len(animals)])
}
