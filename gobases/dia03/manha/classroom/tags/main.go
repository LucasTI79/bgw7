package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name       string  `json:"name"`
	Gender     string  `json:"gender"`
	Age        int     `json:"age"`
	Profession string  `json:"profession"`
	Weight     float64 `json:"weight"`
}

func main() {
	p1 := Person{"Celeste", "Woman", 34, "Engineer", 65.5}

	jsonData, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))

	var p2 Person
	err = json.Unmarshal(jsonData, &p2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", p2)
}
