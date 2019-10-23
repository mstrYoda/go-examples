package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4}

	s2 := []int{5, 6, 7}

	partOfS1 := s1[1:3]

	fmt.Printf("s1 : %v, len %d, cap %d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2 : %v, len %d, cap %d\n", s2, len(s2), cap(s2))
	fmt.Printf("partOfS1 : %v, len %d, cap %d\n", partOfS1, len(partOfS1), cap(partOfS1))

	fmt.Println(partOfS1[3])

	// output :
	/*
		s1 : [1 2 3 4], len 4, cap 4
		s2 : [5 6 7], len 3, cap 3
		partOfS1 : [2 3], len 2, cap 3

		partOfS1 cap = 3 because of it is a pointer to array s1 from element index 1 and length of s1 is 3 from element index 1 to end
	*/
}
