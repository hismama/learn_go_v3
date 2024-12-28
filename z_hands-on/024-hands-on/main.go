package main

import (
	"fmt"
	"math/rand/v2"
)

func init() {
	fmt.Println("My program is initializing...")
}

func main() {
	x := rand.IntN(251)
	fmt.Printf("Name is 'x': %v \t %T\n", x, x)

	if x <= 100 {
		fmt.Println("Between 0 and 100")
	} else if x <=200 {
		fmt.Println("Between 101 and 200")
	} else {
		fmt.Println("Between 201 and 250")
	}

	switch {
	case x <= 100:
		fmt.Println("Between 0 and 100")
	case x <= 200:
		fmt.Println("Between 101 and 200")
	default:
		fmt.Println("Between 201 and 250")
	}
}
