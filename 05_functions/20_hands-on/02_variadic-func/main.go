package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(foo(xi...))
	fmt.Println(bar(xi))

}

func foo(n ...int) int {
	sum := 0
	for _, i := range n {
		sum += i
	}
	return sum
}

func bar(xi []int) int {
	sum := 0
	for _, i := range xi {
		sum += i
	}
	return sum
}

/*
create a func with the identifier foo that
takes in a variadic parameter of type int
pass in a value of type []int into your func (unfurl the []int)
returns the sum of all values of type int passed in

create a func with the identifier bar that
takes in a parameter of type []int
returns the sum of all values of type int passed in

*/
