package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("Hello from Goroutine")
		ch <- 1
	}()
	fmt.Println("Hello from Main")
	//<-ch
	close(ch)
}
