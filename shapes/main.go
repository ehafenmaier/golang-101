package main

import "fmt"

// Shape interface
type shape interface {
	getArea() float64
}

// Shape structs
type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base float64
}

func main () {
	sqr := square{ sideLength: 12}
	tri := triangle{ base: 12, height: 15}

	fmt.Println("Area of a square with a side length of", sqr.sideLength)
	printArea(sqr)

	fmt.Println("Area of a triangle with a base of", tri.base, "and a height of", tri.height)
	printArea(tri)
}

// Struct receiver functions
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return t.base * t.height
}

// Interface functions
func printArea(s shape) {
	fmt.Println("The area of the shape is", s.getArea())
}