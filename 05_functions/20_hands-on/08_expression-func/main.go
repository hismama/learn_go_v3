package main

import "fmt"

func main() {

	xf := func(i int) int {
		return i * 3
	}

	answer := xf(5)
	fmt.Println(answer)

	answer2 := xf(4)
	fmt.Println(answer2)

}
