package main

import "fmt"

// Defer functions run LIFO - last in first out

func main() {
	defer fmt.Println(foo3())
	defer fmt.Println(foo1())
	fmt.Println(foo2())
	defer fmt.Println(foo1())
	defer fmt.Println(foo3())
	defer fmt.Println(foo1())
}

func foo1() int {
	return 1
}

func foo2() int {
	return 2
}

func foo3() int {
	return 3
}
