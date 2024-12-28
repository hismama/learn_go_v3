package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	for n := 0; n < 100; n++ {
		fmt.Println(n)
	}
	
	for n:=0; n < 100; n++ {
		looping_test()
	}
}

func looping_test() {
	x := rand.IntN(10)
	y := rand.IntN(10)
	fmt.Printf("x and y are %v and %v \n", x, y)

	switch {
	case x < 4 && y < 4:
		fmt.Println("x and y are both less than 4")
	case x > 6 && y > 6:
		fmt.Println("x and y are both greater than 6")
	case x >= 4 && x <= 6:
		fmt.Println("x is greater than or equal to 4 and less than or equal to 6")
	case y != 5:
		fmt.Println("y is not 5")
	default:
		fmt.Println("No other case met")
	}
}
