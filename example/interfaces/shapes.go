package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float64
}

type Rectangle struct {
	width float64
	height float64
}

type Square struct {
	width float64
}

type Circle struct {
	radius float64 
}

func (rectan Rectangle) Area() float64 {
	return rectan.width * rectan.height
}

func (sqr Square) Area() float64 {
	return math.Pow(sqr.width, 2.0)
}

func (circle Circle) Area() float64 {
	return math.Pow(circle.radius, 2.0) * math.Pi
}

func computeArea(shape Shaper) {
	fmt.Printf("Area of %+v = %.3f\n", shape, shape.Area())
}