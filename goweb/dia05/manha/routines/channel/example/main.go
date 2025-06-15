package main

import "fmt"

func Sum(num1, num2 int, result chan<- int) {
	result <- num1 + num2
}

func main() {
	channel := make(chan int)
	// var channel chan int

	go Sum(1, 2, channel)

	x := <-channel

	fmt.Println("result:", x)
}
