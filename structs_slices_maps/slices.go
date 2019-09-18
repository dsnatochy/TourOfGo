package main

import "fmt"

// slices are more common than arrays
// a slice is a dynamically-sized, flexible view into the elements of an array
// updating the slice, updates the original array
func main() {
	primes := [6]int{2,3,5,7, 11,13}

	var s[]int = primes[1:4] // action [n,m)
	s[1] = 55  // <- updates the original array
	fmt.Println(s) // [3,55,7]
	fmt.Println(primes) // [2 3 55 7 11 13]
}
