package main

import (
	"fmt"
)

func test(c chan int) {
	fmt.Println("test listening")
	for {
		x := <-c
		fmt.Println("test received : ", x)
	}
}

func main() {
	ch := make(chan int)

	go test(ch)

	ch <- 1
	ch <- 2
	ch <- 3

	// output :
	/*
		test listening
		test received :  1
		test received :  2
		test received :  3
	*/
}
