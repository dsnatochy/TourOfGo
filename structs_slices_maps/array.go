package main

import "fmt"

func main() {
	// type [n]T is an array of n values of type T.
	// array's len is part of its type, so arrays cannot be resized.
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1]) // Hello World
	fmt.Println(a)          // [Hello World]

	primes := [6]int{2,3,5,6,11, 13} // [2 3 5 6 11 13]
	fmt.Println(primes)
}
