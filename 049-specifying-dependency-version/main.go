package main

import (
	"fmt"
	"github.com/hismama/puppy"
)

func main() {
	puppy.From10()
	puppy.From11()
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := puppy.BigBark()
	s4 := puppy.BigBarks()
	fmt.Println(s3)
	fmt.Println(s4)
}