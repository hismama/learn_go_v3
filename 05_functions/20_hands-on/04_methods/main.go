package main

import "fmt"

type person struct {
	first string
	age   int
}

func main() {
	p := person{
		first: "Reid",
		age:   36,
	}
	p.speak()
}

func (p person) speak() {
	fmt.Println("My name is", p.first, "and I am", p.age, "years old.")
	fmt.Printf("My name is %v and I am %v years old.", p.first, p.age)
}
