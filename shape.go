package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (t *triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s *square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

func main() {
	s1 := square{10}
	s2 := triangle{10, 10}

	printArea(&s1)
	printArea(&s2)

	fmt.Println(s1.getArea()) // automatic conversion for pointer receiver
}
