package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("Go routines:", runtime.NumGoroutine())
	var wg sync.WaitGroup

	wg.Add(2)

	go func(s string) {
		fmt.Println("Hello World " + s)
		wg.Done()
	}("1")

	go func(s string) {
		fmt.Println("Hello World " + s)
		wg.Done()
	}("2")
	fmt.Println("Go routines:", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("Go routines:", runtime.NumGoroutine())
}
