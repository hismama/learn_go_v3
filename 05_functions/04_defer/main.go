package main

import "fmt"

func main() {
	defer foo()
	// defer statement execution deferred to the moment surrounding function returns
	bar()
}

// func (r receiver) identifier(p parameter(s)) (r return(s)) { code }

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
