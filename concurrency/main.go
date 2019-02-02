package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(" First Goroutine is Sleeping ")
		time.Sleep(1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(" Second goroutine is sleeping ")
	}()
	wg.Wait()
	fmt.Println("All goroutines complete")
}
