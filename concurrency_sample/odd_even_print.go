package main

import "fmt"
import "sync"

func main() {
	oddch := make(chan int)
	evench := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go odd(oddch, evench, &wg)
	go even(oddch, evench, &wg)
	oddch <- 1
	wg.Wait()
}

func odd(ch1, ch2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch1 {
		fmt.Println(i)
		ch2 <- i + 1
	}
	close(ch2)
}

func even(ch1, ch2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch2 {
		fmt.Println(i)
		if i == 10 {
			break
		}
		ch1 <- i + 1
	}
	close(ch1)
}
