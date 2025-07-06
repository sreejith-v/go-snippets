package main

import "fmt"

func processTask(ch chan int, worker int) {
	for n := range ch {
		fmt.Println("Processing by worker:", worker, " n= ", n, " square: ", n*n)
	}
}

func assignWorkers(ch chan int, workers int) {
	for i := 0; i < workers; i++ {
		go processTask(ch, i+1)
	}
}

func main() {
	limit := 50
	workers := 10

	ch := make(chan int)
	assignWorkers(ch, workers)

	for i := 0; i < limit; i++ {
		ch <- i
	}
	close(ch)
}
