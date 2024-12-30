package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I am walking.")
}

func (d *dog) run() {
	d.first = "Rover" // mutated
	fmt.Println("My name is", d.first, "and I am running.")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
}

func main() {
	d1 := dog{"Henry"}
	d1.walk()
	d1.run()
	d1.walk()
	// youngRun(d1)

	d2 := &dog{"Padget"}
	d2.walk()
	d2.run()
	d2.walk()
	youngRun(d2)

}
