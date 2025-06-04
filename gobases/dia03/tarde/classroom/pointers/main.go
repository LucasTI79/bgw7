package main

import "fmt"

// Increase receives an integer pointer
func Increase(v *int) {
	// We dereference the variable v to obtain
	// its value and increment it by 1
	*v++
}

func main() {
	var v int = 19
	// Increase function receives a pointer
	// we use the address operator &
	// to pass the memory address of v
	Increase(&v)
	fmt.Println("The value of v now reads:", v)
}
