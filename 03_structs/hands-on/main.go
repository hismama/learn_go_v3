package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   37,
	}

	// Anonymous Struct
	p2 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   37,
	}

	fmt.Println(p1)
	fmt.Println(p2)
}
