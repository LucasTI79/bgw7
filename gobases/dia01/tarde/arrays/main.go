package main

import "fmt"

func main() {
	var a [2]string

	a[0] = "Hello"
	a[1] = "World"
	// a[2] = "panico"

	fmt.Println(a[0], a[1])
}
