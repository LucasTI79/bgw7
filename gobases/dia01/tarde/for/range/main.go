package main

import "fmt"

// string
// int
// float
// bool

func main() {
	fruits := []string{"apple", "banana", "pear"}
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}
	fmt.Println(fruits)
}
