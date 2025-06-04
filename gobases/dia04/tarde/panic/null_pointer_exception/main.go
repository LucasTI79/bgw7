package main

import "fmt"

type Dog struct {
	Name string
}

func (s *Dog) WoofWoof() {
	fmt.Println(s.Name, "Goes woof woof")
}

func main() {
	s := &Dog{"Sammy"}
	s = nil
	s.WoofWoof()
}
