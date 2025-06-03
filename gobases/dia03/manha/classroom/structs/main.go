package main

import "fmt"

type Person struct {
	Name       string
	Gender     string
	Age        int
	Profession string
	Weight     float64
	Likes      Preferences
}

type Preferences struct {
	Foods  string
	Movies string
	Series string
	Animes string
	Sports string
}

func main() {
	p1 := Person{"Celeste", "Woman", 34, "Engineer", 65.5, Preferences{"chicken", "titanic", "", "", ""}}
	p2 := Person{
		Name: "Fulano",
		Likes: Preferences{
			Movies: "WALL-E",
		},
	}

	fmt.Printf("%+v", p1)
	fmt.Println(p2)
}
