package main

import "fmt"

func main() {
for1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; i++ {
			if i == 5 && j == 4 {
				break for1
			}
			fmt.Print(j, i)
		}
	}
}
