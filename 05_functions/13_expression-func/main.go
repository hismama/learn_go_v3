package main

import "fmt"

func main() {
	foo()

	x := func() {
		fmt.Println("Anonymous foo")
	}

	y := func(s string) {
		fmt.Println("This is an anonymous func showing my name", s)
	}

	x()
	y("Reid")

}

func foo() {
	fmt.Println("foo ran")
}

// A named function
// func (r receiver) identifier(p parameters) (r returns) { code }

// An anonymous func
// func(p parameters) (r returns) { code }
