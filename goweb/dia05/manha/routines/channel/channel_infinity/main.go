package main

import (
	"fmt"
	"time"
)

func processMessage(queue chan int) {
	func() {
		i := 0
		for {
			time.Sleep(time.Second)
			queue <- i
			i++
		}
	}()
}

func main() {
	queue := make(chan int)

	go processMessage(queue)

	for x := range queue {
		fmt.Println(x)
	}
}
