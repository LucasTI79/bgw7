package main

import (
	"fmt"
	"time"
)

func worker(workerId int, msg chan int, results chan int) {
	for res := range msg {
		fmt.Println("Worker:", workerId, "msg:", res)
		results <- res
	}
	fmt.Println("Worker", workerId, "finalizou")
}

func main() {
	msg := make(chan int)
	results := make(chan int)

	// Inicia 11 workers
	for i := 1; i <= 11; i++ {
		go worker(i, msg, results)
	}

	// Envia 10 mensagens
	for i := 0; i < 10; i++ {
		msg <- i
	}

	// Espera um pouco para os workers terminarem antes de sair da main
	time.Sleep(3 * time.Second)

	close(msg)
	close(results)

	for result := range results {
		fmt.Println("result", result)
	}

	fmt.Println("Main finalizou")
}
