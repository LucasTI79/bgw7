package main

import (
	"encoding/json"
	"os"
)

func main() {
	// It is necessary to create our type Encode
	// for this the NewEncoder function is called
	// this receives a streaming as a parameter
	// we use one of the standard streams offered by the OS Stdout pkg
	// stdout generates a stream to a file that is printed to the console.

	myEncoder := json.NewEncoder(os.Stdout)

	// prepare the information you want to send in json format to the streaming
	type MyData struct {
		ProductID string
		Price     float64
	}

	data := MyData{
		ProductID: "XASD",
		Price:     25.50,
	}

	// the Encode method is invoked.
	// internally this method makes a kind of marshall and writes it to the stream
	myEncoder.Encode(data)
}
