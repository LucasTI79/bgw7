package main

import (
	"fmt"
	"time"
)

func Sum(num1, num2 int) {
	fmt.Println("result:", num1+num2)
}

func Sub(num1, num2 int) {
	fmt.Println("result:", num1-num2)
}

func main() {
	go Sum(1, 2)

	time.Sleep(1 * time.Second)
	fmt.Println("program finished")
}

// CPU // Calculos matematicos
