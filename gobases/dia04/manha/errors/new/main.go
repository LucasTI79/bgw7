package main

import (
	"errors"
	"fmt"
)

// type MyError struct {
// 	Message string
// }

// func (e *MyError) Error() string {
// 	return e.Message
// }

func validateStatusCode(code int) error {
	if code > 399 {
		return errors.New("unexpected http status code")
	}
	return nil
}

func main() {
	statusCode := 404
	if err := validateStatusCode(statusCode); err != nil {
		fmt.Printf("http request failed: %v", err)
		return
	}
	fmt.Println("the program ended successfully")
}
