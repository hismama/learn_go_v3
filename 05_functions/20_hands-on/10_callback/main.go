package main

import "fmt"

func main() {
	x := operate(4, 6, multiply)
	fmt.Println(x)

	y := operate(8, 3, divide)
	fmt.Println(y)

}

func operate(x, y int, f func(int, int) float64) float64 {
	return f(x, y)
}

func multiply(x, y int) float64 {
	return float64(x * y)
}

func divide(x, y int) float64 {
	if y == 0 {
		return 0
	}
	return float64(x) / float64(y)
}
