package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("starting... ")
	_, err := os.Open("no-file.txt")
	// defer file.Close()

	if err != nil {
		panic(err)
	}
	fmt.Println("end")
}
