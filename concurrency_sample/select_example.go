package main

import "fmt"

func directChannelPrint(ch1, ch2 chan int) {
	go func() {
		ch1 <- 1
		ch2 <- 2
		close(ch1)
		close(ch2)
	}()
	// if the order for above is changed, func will break

	one := <-ch1
	two := <-ch2
	fmt.Println(one, two)
}

func selectChannelPrint(ch1, ch2 chan int) {
	go func() {
		ch1 <- 1
		ch2 <- 2
		close(ch1)
		close(ch2)
	}()
	// will print in exact order of receiving, will not break if order changed

	for i := 0; i < 2; i++ { // need loop to print both cases
		select {
		case one := <-ch1:
			fmt.Println(one)
		case two := <-ch2:
			fmt.Println(two)
		}
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	ch3 := make(chan int)
	ch4 := make(chan int)

	directChannelPrint(ch1, ch2)
	selectChannelPrint(ch3, ch4)
}
