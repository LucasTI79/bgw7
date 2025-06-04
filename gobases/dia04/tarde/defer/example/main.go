package main

import (
	"fmt"
)

func main() {
	//apply “defer” to the invocation of an anonymous function
	defer func() {
		fmt.Println("This function is executed despite a panic occurring")
	}()
	//create a panic with a message that it occurred
	// os.Exit(2)
	panic("panic occurred!!!")
}
