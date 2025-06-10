package main

import (
	"fmt"
	"sync"
)

func printOdd(limit int, res chan string, wg *sync.WaitGroup) {
	for i := 1; i <= limit; i++ {
		if i%2 != 0 {
			s := fmt.Sprintf("Odd number: %v", i)
			res <- s
		}
	}
	wg.Done()
}

func printEven(limit int, res chan string, wg *sync.WaitGroup) {
	for i := 1; i <= limit; i++ {
		if i%2 == 0 {
			s := fmt.Sprintf("Even number: %v", i)
			res <- s
		}
	}
	wg.Done()
}

func printOddNew(limit int, oddRes chan int, evenRes chan int, res chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range evenRes {
		if i > limit-1 {
			close(oddRes)
			break
		}
		s := fmt.Sprintf("Odd number: %v", i+1)
		res <- s
		oddRes <- i + 1
	}
}

func printEvenNew(limit int, oddRes chan int, evenRes chan int, res chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range oddRes {
		if i > limit-1 { //todo: check edge case
			close(evenRes)
			break
		}
		s := fmt.Sprintf("Even number: %v", i+1)
		res <- s
		evenRes <- i + 1
	}
}

func main() {
	limit := 10
	channel := make(chan string)
	wg := new(sync.WaitGroup)
	oddRes := make(chan int)
	evenRes := make(chan int)

	//wg.Add(2)
	//go printOdd(limit, channel, wg)
	//go printEven(limit, channel, wg)
	//
	//go func() {
	//	wg.Wait()
	//	close(channel)
	//}()
	//
	//for i := range channel {
	//	fmt.Println(i)
	//}

	wg.Add(2)
	go func() {
		evenRes <- 0
	}()

	go printOddNew(limit, oddRes, evenRes, channel, wg)
	go printEvenNew(limit, oddRes, evenRes, channel, wg)

	go func() {
		wg.Wait()
		close(channel)
	}()

	for i := range channel {
		fmt.Println(i)
	}
}
