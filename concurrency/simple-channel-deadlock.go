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

	ch <- 1
	ch <- 2

	fmt.Println("main end")
}
