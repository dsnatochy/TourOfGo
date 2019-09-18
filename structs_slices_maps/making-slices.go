package main

import "fmt"

func main() {
	// "make" function allocates a zeroed array and returns a slice that refers to that array
	a := make([]int, 5)
	printSlice2("a", a)  // a  len=5 cap=5 [0 0 0 0 0]

	// to specify a capacity, pass a third argument to "make" (5)
	b := make([]int, 0, 5)
	// b := make([]int, 0, 6)  // passed compile check, but b[5] will generate out of range error
	printSlice2("b", b) // b  len=0 cap=5 []

	c := b[:2]
	printSlice2("c", c) // c  len=2 cap=5 [0 0]

	d := c[2:5]
	printSlice2("d", d) // d  len=3 cap=3 [0 0 0]
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s  len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
