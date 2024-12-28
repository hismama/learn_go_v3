package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Printf("shape_area: %f\n", s.area())
	fmt.Printf("shape_area: %v\n", s.area())
}

func main() {
	sq := square{length: 5}
	cir := circle{radius: 5}

	info(sq)
	info(cir)

}
