package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) speak() {
	fmt.Printf("%s is %d years old.\n", p.name, p.age)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		name: "Bob",
		age:  20,
	}
	p1 := &person{
		name: "James",
		age:  34,
	}
	saySomething(p1)
	//saySomething(p)
	saySomething(&p)
	p.speak()

	//	Other ways to make a pointer of type person '*person'
	/*
		   p2 := new(person)
		   p2.name = "Alice"
		   p2.age = 25

			func newPerson(name string, age int) *person {
				return &person{name: name, age: age}
			}
			p5 := newPerson("Eve", 30)

	*/

}
