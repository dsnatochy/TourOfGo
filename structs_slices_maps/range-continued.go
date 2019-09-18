package main

import "fmt"

func main() {
	pow := make([]int, 10) // [0 0 0 0 0 0 0 0 0 0]

	// if you only want the index, you can omit the second var
	// for i := range pow
	for i:= range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}

	// index or value can be skipped by assigning to _
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
	/*
		1
		2
		4
		8
		16
		32
		64
		128
		256
		512
	 */
}
