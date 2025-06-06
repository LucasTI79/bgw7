package main

import (
	"encoding/json"
	"io"
	"log"
	"strings"
)

func main() {
	const jsonStream = `
    {"ProductID": "AXW123", "Price": 25.50}
    {"ProductID": "NLBR17", "Price": 357.58}
    {"ProductID": "KNUB82", "Price": 150}
    `
	// It is necessary to create our type Decode, for this the NewDecoder function is called
	// this receives a streaming as a parameter
	// A jsonStream variable is created and the NewReader method of the strings pkg is used
	// NewReader generates a streaming for the text string it receives.

	myStreaming := strings.NewReader(jsonStream)
	myDecoder := json.NewDecoder(myStreaming)

	type MyData struct {
		ProductID string
		Price     float64
	}
	// streaming behaves so that each line in the jsonStrem text is streamed separately
	// we go through all the data transmitted in the streaming until the end of the text is reached
	for {
		// the variable on which the data is going to be written is created
		var data MyData
		// the decode method is invoked
		// Decode is responsible for reading the data transmitted by the streaming and transforming it from json to our interface
		if err := myDecoder.Decode(&data); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		// The received data is printed
		log.Println("Data:", data.ProductID, data.Price)
	}

}
