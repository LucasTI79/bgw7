package main

import (
	"fmt"
	"time"
)

type MathOperation interface {
	// Do is a method that receives two integers and returns an integer
	Do(int, int) int
}

// Sum is a struct that implements the MathOperation interface
type Sum struct{}

// Do is a method that receives two integers and returns an integer
func (s Sum) Do(n1, n2 int) int {
	return n1 + n2
}

type ShowTime struct {
	// MathOperation represents the interface that will be decorated
	MathOperation
}

// Do is a method that receives two integers and returns an integer
func (s ShowTime) Do(n1, n2 int) int {
	// display info about the time
	fmt.Println("time:", time.Now())

	// call the function
	return s.MathOperation.Do(n1, n2)
}

func main() {
	fn := Sum{}

	st := ShowTime{MathOperation: fn}

	fmt.Println("result:", st.Do(1, 2))

}
