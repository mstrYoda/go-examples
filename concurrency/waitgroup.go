package main

import (
	"fmt"
	"sync"
)

func run1(wg *sync.WaitGroup) {
	fmt.Println("run1")
	wg.Done()
}

func run2(wg *sync.WaitGroup) {
	fmt.Println("run2")

	// we call we are done
	wg.Done()
}

func main() {

	// wait group provides to wait all goroutines to end before exit main goroutine

	wg := sync.WaitGroup{}

	// here we give an integer value to how many done are we call
	wg.Add(2)

	go run1(&wg)
	go run2(&wg)

	// we wait for all goroutines to complete
	wg.Wait()
}
