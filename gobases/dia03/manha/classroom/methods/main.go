package main

import (
	"fmt"

	"google.com/bgw7/methods/geometry"
)

func main() {
	circle1 := geometry.NewCircle(2)

	fmt.Println("area: ", circle1.Area())
	fmt.Println("perimetro: ", circle1.Perimeter())

	circle1.SetRadius(3)

	fmt.Println("area: ", circle1.Area())
	fmt.Println("perimetro: ", circle1.Perimeter())
}
