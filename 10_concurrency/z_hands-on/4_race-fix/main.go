package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	var wg sync.WaitGroup
	var mu sync.Mutex

	counter := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			fmt.Println(counter)
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("end value:", counter)
	fmt.Println("elapsed:", elapsed)
}
