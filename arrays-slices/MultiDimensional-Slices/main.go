package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	jm := []string{"Jenny", "Moneypenny", "Guiness", "Wolverine Tracks"}

	xxs := [][]string{jb, jm}  //storing string slices into another slice
	fmt.Println(xxs)

	// Setting one slice equal to another, both will point to the same underlying array and be equal
	// Use make then copy if you want to keep separate
}