package main

import (
	"fmt"
)

func produce(ch chan int) {
	ch <- 1
}

func main() {

	ch := make(chan int)
	ch2 := make(chan int)

	go produce(ch)
	go produce(ch2)

	end := 0
	for {
		fmt.Println("for")
		select {
		case x := <-ch:
			fmt.Println("value from ch : ", x)
			end++
		case x := <-ch2:
			fmt.Println("value from ch2 : ", x)
			end++
		}

		if end == 2 {
			break
		}
	}

	close(ch)
	close(ch2)
}
