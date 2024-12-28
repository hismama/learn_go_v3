package main

import "fmt"

func main() {

	// 42
	xa := [5]int{0, 1, 2, 3, 4}
	for _, i := range xa {
		fmt.Printf("%v  %T\n", i, i)
	}

	fmt.Println("----------")

	// 43
	xi := []int{}
	for i := 42; i < 52; i++ {
		xi = append(xi, i)
	}
	for i, v := range xi {
		fmt.Printf("%v - %v as %T\n", i, v, v)
	}
	fmt.Println("----------")

	// 44
	a1 := xi[:5]
	a2 := xi[5:]
	a3 := xi[2:7]
	a4 := xi[1:6]
	fmt.Printf("%v\n%v\n%v\n%v\n\n", a1, a2, a3, a4)

	fmt.Println("----------")

	// 45
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)

	fmt.Println("----------")

	// 46
	x = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x1 := x[:3]
	x2 := x[6:]
	x = append(x1, x2...)
	fmt.Println(x)

	fmt.Println("----------")

	// 47
	states := make([]string, 0, 50)
	states = append(states, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, `
		Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `
		Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `
		Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
		` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
		` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, `Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	fmt.Printf("Len: %v, cap: %v , %#v\n", len(states), cap(states), states)
	for i:=0; i<len(states); i++{
		fmt.Println(states[i], i)
	}

	
	fmt.Println("----------")

	// 48
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "I'm 008"}
	xss := [][]string{jb, mp}
	
	for i , r := range xss{
		fmt.Println(i, r)
		for a, b := range r {
			fmt.Println(a, b)
		}
		fmt.Println("")
	}
}
