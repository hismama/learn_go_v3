package main

import "fmt"

type person struct {
	first string
}

func main() {

	p1 := person{
		first: "James",
	}
	fmt.Println(p1.first)
	p1.first = valueChange("James")
	fmt.Println(p1.first)
	pointerChange(&p1)
	fmt.Println(p1.first)

}

func valueChange(s string) string {
	return s + "world"
}

func pointerChange(p *person) {
	p.first = "Hello"
}

// Todd's VALUE adjust
func changeName(p person, s string) person {
	p.first = s
	return p
}

// Todd's POINTER adjust
func changeNamePtr(p *person, s string) {
	p.first = s
}
