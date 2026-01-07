package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func genRandomNums(ch chan<- int, wg *sync.WaitGroup) {
	num := rand.Intn(500)

	ch <- num
	wg.Done()
}

func main() {
	limit := 10
	ch := make(chan int, limit)
	wg := sync.WaitGroup{}

	for i := 0; i < limit; i++ {
		wg.Add(1)
		go genRandomNums(ch, &wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for num := range ch {
		fmt.Println(num)
	}
}
