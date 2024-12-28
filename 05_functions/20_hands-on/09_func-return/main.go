package main

import "fmt"

func main() {
	x := foo()
	y := bar()
	fmt.Println(x)
	fmt.Println(y())
	fmt.Println(x, y())
}

func foo() string {
	return "hello"
}

func bar() func() string {
	return inside
}

func inside() string {
	return "world"
}
