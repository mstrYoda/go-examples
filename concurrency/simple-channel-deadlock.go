package main

import "fmt"

func listen(ch chan int) {
	defer fmt.Println("listen goroutine end")
	fmt.Println("started listen to channel")
	x := <-ch
	fmt.Println(x)
}

func main() {

	ch := make(chan int)

	go listen(ch)

	// here go scheduler blocks main goroutine to schedule another goroutine to read from channel
	// when it first sends data to channel, listen() goroutine gets value from channel end it ends
	// when we send another data, it will get "fatal error: all goroutines are asleep - deadlock!" error
	// because there is not any other listener for the channel
	ch <- 1
	ch <- 2

	fmt.Println("main end")
}
