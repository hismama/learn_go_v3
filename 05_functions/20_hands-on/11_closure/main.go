package main

import (
	"fmt"
	"math"
)

func main() {
	x := powinator(2)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}

func powinator(n float64) func() float64 {
	var x float64
	return func() float64 {
		x++
		//fmt.Println("this is x:", x)
		//fmt.Println("this is n:", n)
		return math.Pow(n, x)
	}
}
