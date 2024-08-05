package main

import (
	"fmt"
	"math"
)

type retangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

type shape interface {
	getArea() float64
}

func (retangle retangle) getArea() float64 {
	return retangle.width * retangle.height
}

func (circle circle) getArea() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}

func writeArea(shape shape) {
	fmt.Printf("The shape of the form has an area of %.2f\n", shape.getArea())
}

func main() {
	retangle1 := retangle{10, 15}
	circle1 := circle{10}

	writeArea(retangle1)
	writeArea(circle1)
}