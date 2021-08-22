package main

import "fmt"

type triangle struct {
	height, base float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func (t triangle) getArea() float64 { 
	return 0.5 * t.height * t.base
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(shp shape) {
	fmt.Printf("Area of %v is: %v\n",shp ,shp.getArea())
}

func main() {
	tri := triangle{height: 10, base: 20}
	sqr := square{sideLength: 20}

	printArea(tri)
	printArea(sqr)
}
