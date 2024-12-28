package main

import "fmt"

func main() {
	//	4!
	fmt.Println(4 * 3 * 2 * 1)
	fmt.Println(factorial(4))
}

func factorial(n int) int {
	fmt.Println("This is n", n)
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
