package main

import "fmt"

func main() {
	arr := [4]int{1, 2, 3, 4}

	s1 := arr[1:3]

	fmt.Println(arr)
	fmt.Println(s1)

	// output:
	/*
		[1 2 3 4]
		[2 3]
	*/

	// we change slice element, and it affects the array
	// remember, slices are pointers to arrays
	s1[0] = 11

	fmt.Println(arr)
	fmt.Println(s1)

	// output:
	/*
		[1 11 3 4]
		[11 3]
	*/
}
