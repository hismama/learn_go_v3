package main

import "fmt"

func main() {
	fmt.Println("f")
	f := incrementor()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 2
	fmt.Println(f()) // 3
	fmt.Println(f()) // 4
	fmt.Println(f()) // 5
	fmt.Println(f()) // 6

	g := incrementor()
	fmt.Println("\ng")
	fmt.Println(g()) // 1
	fmt.Println(g()) // 2
	fmt.Println(g()) // 3
	fmt.Println(g()) // 4
	fmt.Println(g()) // 5
	fmt.Println(g()) // 6

	fmt.Println("\nf")
	fmt.Println(f()) // 7
}

// Closure to have func inside. Powerful technique
func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
