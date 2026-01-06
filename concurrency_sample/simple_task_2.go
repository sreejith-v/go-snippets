package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func routine(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	val := <-ch
	val++
	fmt.Println("Routine: ", val)
	ch <- val
}

func routine2(id int, counter *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	//mu.Lock()
	*counter++
	fmt.Printf("Goroutine %d incremented counter to %d\n", id, *counter)
	//mu.Unlock()
}

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int, 3)
	//log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("Starting concurrent routines")
	startTime1 := time.Now()
	ch <- 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go routine(ch, &wg)
	}
	wg.Wait()
	log.Println("Elapsed time for first part:", time.Since(startTime1))
	close(ch)

	// another solution

	var counter int
	var mu sync.Mutex

	startTime2 := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go routine2(i, &counter, &mu, &wg)
	}

	wg.Wait()
	log.Println("Elapsed time for second part:", time.Since(startTime2))
	fmt.Println("Final counter value:", counter)
}
