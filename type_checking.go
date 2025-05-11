package main

import "fmt"

func main() {
	var t interface{}

	t = 10.0

	fmt.Printf("type of t=%T\n", t)

	a, b := t.(float64)

	fmt.Println(a, b)
}
