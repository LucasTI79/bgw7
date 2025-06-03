package main

import "fmt"

type Engine struct {
	Horsepower int
	Type       string
}
type Chassis struct {
	Material string
}
type Bodywork struct {
	Color string
}

type Vehicle interface {
	Start()
	DisplayInfo()
}

type Car struct {
	Engine
	Chassis
	Bodywork
	NumberOfDoors int
}

func (c Car) Start() {
	fmt.Println("The car is starting with a roar!")
}
func (c Car) DisplayInfo() {
	fmt.Printf(
		"This car has %d doors, a %s body, a %s chassis, and a %d horsepower %s engine.\n",
		c.NumberOfDoors,
		c.Bodywork.Color,
		c.Chassis.Material,
		c.Engine.Horsepower, c.Engine.Type,
	)
}

type Motorcycle struct {
	Engine
	Chassis
	Bodywork
}

func (m Motorcycle) Start() {
	fmt.Println("The motorcycle is starting with a vroom!")
}
func (m Motorcycle) DisplayInfo() {
	fmt.Printf(
		"This motorcycle has a %s body, a %s chassis, and a %d horsepower %s engine.\n",
		m.Bodywork.Color,
		m.Chassis.Material,
		m.Engine.Horsepower,
		m.Engine.Type,
	)
}

func main() {
	myCar := Car{
		Engine:        Engine{Horsepower: 250, Type: "V4"},
		Chassis:       Chassis{Material: "Steel"},
		Bodywork:      Bodywork{Color: "Red"},
		NumberOfDoors: 4,
	}

	myMotorcycle := Motorcycle{
		Engine:   Engine{Horsepower: 150, Type: "V2"},
		Chassis:  Chassis{Material: "Aluminum"},
		Bodywork: Bodywork{Color: "Black"},
	}

	var vehicle Vehicle

	vehicle = myCar

	vehicle.DisplayInfo()

	vehicle = myMotorcycle

	vehicle.DisplayInfo()

}
