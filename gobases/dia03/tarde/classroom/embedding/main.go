package main

import "fmt"

type Author struct {
	FirstName string
	LastName  string
}

func (a *Author) fullName() string {
	return fmt.Sprintf("%s %s", a.FirstName, a.LastName)
}

type Book struct {
	Title   string
	Content string
	Autor   Author
}

func (b *Book) information() {
	fmt.Println("Title: ", b.Title)
	fmt.Println("Content: ", b.Content)
	fmt.Println("Autor: ", b.Autor.fullName())
}

func main() {
	autor := Author{
		"Juan",
		"Lopez",
	}

	book := Book{
		Title:   "Inheritance in Go",
		Content: "Go has composition instead of inheritance",
		Autor:   autor,
	}
	book.information()
}
