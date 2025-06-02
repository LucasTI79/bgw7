package main

import "fmt"

func main() {
	sum := 0
	for {
		sum++
		if sum >= 1000 {
			break
		}
	}
	fmt.Println(sum)
}
