package main

import (
	"fmt"
	"time"
)

func worker(workerId int, msg chan int) {
	for res := range msg {
		fmt.Println("Worker: ", workerId, "msg: ", res)
		time.Sleep(time.Millisecond * 50)
	}
}

func main() {
	msg := make(chan int)

	for i := 0; i <= 10; i++ {
		go worker(i+1, msg)
	}

	// for i := 0; i < 10; i++ {
	// 	msg <- i
	// }
	// time.Sleep(1 * time.Second)

	for i := 0; i < 1_000_000; i++ {
		msg <- i
	}
}
