package main

import (
	"errors"
	"fmt"
)

var err1 = errors.New("error number 1")

func x() error {
	// wrap
	// fmt.Printf("%d %s %f %v %T %+v")
	return fmt.Errorf("extra error information: %w", err1)
}

func main() {
	e := x()
	coincidence := errors.Is(e, err1)
	fmt.Println(coincidence) //print true
}
