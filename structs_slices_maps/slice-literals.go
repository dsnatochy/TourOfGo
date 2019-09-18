package main

import "fmt"

func main() {
	// creates an array literal, then builds a slice that references it
	// note: no length
	q := []int{2,3,5,7,11, 13}  // [2 3 5 7 11 13]
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)  //[true false true true false true]

	s := []struct { // array of structs
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(s[2]) // {5 true}
	fmt.Println(s) // [{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}]


}
