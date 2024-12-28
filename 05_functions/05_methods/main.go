package main

import "fmt"

type person struct {
	first string
}

// Creating a method is done via the "receiver" portion of a func "p person"
func (p person) speak() {
	fmt.Println("I am ", p.first)
}

func main() {
	p1 := person{
		first: "James",
	}
	p2 := person{
		first: "Jenny",
	}

	p1.speak()
	p2.speak()
}
