package main

import (
	"errors"
	"fmt"
)

func SayHello(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name provided")
	}
	return fmt.Sprintf("Hi %s", name), nil
}

func main() {
	name := "Meli"
	greeting, err := SayHello(name)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(greeting)
}
