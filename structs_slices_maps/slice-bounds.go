package main

import "fmt"

func main() {
	s := []int{2,3,5,7,11, 13}

	s = s[:] // [2 3 5 7 11 13]
	fmt.Println(s)

	s = s[1:4] // [3 5 7]
	fmt.Println(s)

	// high or low bounds can be omiited to use their defaults
	// default is "0" for the low bound and the length of the slice for the high bound
	s = s[:2] // [3 5]
	fmt.Println(s)

	s = s[1:] // [5]
	fmt.Println(s)
}
