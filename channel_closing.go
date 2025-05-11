package main

import "fmt"

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	fmt.Println(<-queue)
	close(queue)
	queue <- "three"

	for elem := range queue {
		fmt.Println("in loop")
		fmt.Println(elem)
	}
}
