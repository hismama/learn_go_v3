package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("Anonymous foo")
	}()

	func(s string) {
		fmt.Println("This is an anonymous func showing my name", s)
	}("Reid")

}

func foo() {
	fmt.Println("foo ran")
}

// A named function
// func (r receiver) identifier(p parameters) (r returns) { code }

// An anonymous func
// func(p parameters) (r returns) { code }
