package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read")

	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		panic(err)
	}
}
