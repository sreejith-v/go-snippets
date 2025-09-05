package prime

import (
	"math"
	"sync"
)

func FindPrimesBasic(N int) []int {
	ans := []int{}
	for i := 2; i <= N; i++ {
		if isPrime(i) {
			ans = append(ans, i)
		}
	}
	return ans
}

func FindPrimesOptimized(N int) []int {
	workers := 10
	ch := make(chan int)
	res := make(chan int)
	wg := sync.WaitGroup{}
	// assign workers
	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := range ch {
				if isPrime(n) {
					res <- n
				}
			}
		}()
	}
	// feeder
	go func() {
		for i := 2; i <= N; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		wg.Wait()
		close(res)
	}()
	// collector
	ans := []int{}
	for i := range res {
		ans = append(ans, i)
	}

	return ans
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	limit := int(math.Ceil(math.Sqrt(float64(num))))
	for i := 3; i <= limit; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}
