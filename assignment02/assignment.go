package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println("The area of this shape is: ", s.getArea())
}

func main() {
	s := square{sideLength: 10}

	printArea(s)
}
