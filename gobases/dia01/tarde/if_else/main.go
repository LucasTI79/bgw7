package main

import "fmt"

func main() {
	var salary float64 = 4000
	if salary <= 3000 {
		fmt.Println("This person must pay taxes")
	} else if salary <= 4000 {
		fmt.Printf("They must pay $%4.2f of their salary\n", (salary/100)*10)
	} else {
		fmt.Printf("They must pay $%4.2f of their salary\n", (salary/100)*15)
	}
}
