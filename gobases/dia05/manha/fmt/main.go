package main

import "fmt"

func main() {
	name, age := "Lucas", 22
	fmt.Print("Hello, my name is ", name, " i'm ", age, "years old\n")

	fmt.Println("Hello, my name is ", name, " i'm ", age, "years old")

	fmt.Printf("Hello, my name is %s i'm %d years old\n", name, age)

	text := fmt.Sprint("Hello, my name is ", name, " i'm ", age, "years old\n")

	text = fmt.Sprintf("Hello, my name is %s i'm %d years old\n", name, age)

	fmt.Print(text)

}
