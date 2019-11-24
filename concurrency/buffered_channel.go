package main

import (
	"fmt"
	"time"
)

func test(c chan int) {
	fmt.Println("test listening")
	time.Sleep(5 * time.Second) // the purpose of time sleep is to show channel blocked after third element
	for {
		select {
		case x := <-c:
			fmt.Println("test received : ", x)
		}
	}
}

func main() {
	ch := make(chan int, 3)

	go test(ch)

	ch <- 1
	fmt.Println("sent 1")
	ch <- 2
	fmt.Println("sent 2")
	ch <- 3
	fmt.Println("sent 3")
	ch <- 4 // blocked here
	fmt.Println("sent 4")

	// output :
	/*
		sent 1
		sent 2
		sent 3
		test listening
		(NOTE: below could change according to go-scheduler)

		test received :  1
		test received :  2
		sent 4
		test received :  3
		test received :  4

		because we used buffered channel, it prevents to send 4th element after third. we can send it only after channel consumed by receiver
	*/
}
