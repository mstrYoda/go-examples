package main

import "fmt"

func produce(ch chan int) {
	defer fmt.Println("produce goroutine end")
	ch <- 1 // if we don't do this we faced with deadlock error because main is blocked with select
}

func main() {

	ch := make(chan int)
	go produce(ch)

	select {
	case x := <-ch:
		fmt.Println("value from ch : ", x)
	}

	fmt.Println("main end")
}
