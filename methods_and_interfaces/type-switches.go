package main

import "fmt"


// A "type switch" is a construct that permits several type assertions in series

//

func do(i interface{}) {
	switch v := i.(type){   // new "type" keyword
	case int:
		fmt.Printf("Twice %v is %v\n", v , v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21) 		// Twice 21 is 42
	do("hello") 	// "hello" is 5 bytes long
	do(true)  	// I don't know about type bool!
}