package main

import "fmt"

func getSecond[second int | float64](a any, b second) second {
	return b
}

func main() {
	second := getSecond(5, 10.5)
	fmt.Println(second)
	fmt.Println(getSecond(5, 1))
}
