package main

import "fmt"

func generateFibonacci(n int, ch, res chan int) {
	num1, num2 := 1, 1
	for i := 1; i <= n; i++ {
		res <- num1
		num1, num2 = num2, num1+num2
	}
	close(res)
}

func main() {
	limit := 10
	ch := make(chan int, 2)
	res := make(chan int)
	go generateFibonacci(limit, ch, res)

	for i := range res {
		fmt.Println(i)
	}
	close(ch)
}
