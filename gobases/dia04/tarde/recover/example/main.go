package main

import "fmt"

func isPair(num int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	// if err := recover(); err != nil {

	// }

	if (num % 2) != 0 {
		panic("Not an even number")
	}

	// se o recover estivesse aqui, ele não seria executado
	// recover()

	fmt.Println(num, "is an even number!")
}

func main() {
	num := 3
	isPair(num)
	fmt.Println("execution completed")
}
