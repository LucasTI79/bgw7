package main

import (
	"errors"
	"fmt"
)

// we define a type struct
type myError struct {
	msg string
	x   string
}

// we make our type struct implement the Error() method
func (e *myError) Error() string {
	return fmt.Sprintf("An error occurred: %s, %s", e.msg, e.x)
}

func main() {
	var targetError *myError

	_, err := GetUsers()

	isMyError := errors.As(err, &targetError) // compares errors

	if isMyError {
		fmt.Println(targetError.Error())
		return
	}

	fmt.Println(isMyError) //prints true because the errors match
}

func GetUsers() (any, error) {
	return nil, &myError{"new error", "404"}
}
