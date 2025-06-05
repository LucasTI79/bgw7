package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read")

	b, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
