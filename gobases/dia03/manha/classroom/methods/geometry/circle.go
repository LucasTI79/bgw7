package geometry

import (
	"math"
)

type Circle struct {
	radius float64
}

func NewCircle(radius float64) Circle {
	return Circle{radius}
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// encapsulamento
func (c *Circle) SetRadius(newRadius float64) {
	c.radius = newRadius
	// fmt.Printf("%+v", c)
}

// quando vamos modificar algum campo da struct, usamos ponteiros
// quando vamos apenas visualizar os valores, ou retornar eles, não é necessário
