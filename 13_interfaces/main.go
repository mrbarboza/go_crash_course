package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x      float64
	y      float64
	radius float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (rectangle Rectangle) area() float64 {
	return rectangle.width * rectangle.height
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	circle := Circle{
		x:      0,
		y:      0,
		radius: 5,
	}

	rectangle := Rectangle{
		width:  10,
		height: 5,
	}

	fmt.Printf("CIRCLE Area - %f\n", getArea(circle))
	fmt.Printf("RECTANGLE Area - %f\n", getArea(rectangle))
}
