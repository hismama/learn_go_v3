package main

import "fmt"

func main() {
	x := 3
	func(i int) {
		fmt.Println(x * i)
	}(5)
}
