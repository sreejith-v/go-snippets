package main

import (
	"cmp"
	"fmt"
)

func maxc[T cmp.Ordered](x, y T) T {
	res := cmp.Compare(x, y)
	if res == 1 {
		return x
	}
	return y
}

func main() {
	maxInt := maxc(5, 10)
	fmt.Println("Max integer:", maxInt)

	maxFloat := maxc(3.14, 2.71)
	fmt.Println("Max float:", maxFloat)

	maxString := maxc("hello", "world")
	fmt.Println("Max string:", maxString)
}
