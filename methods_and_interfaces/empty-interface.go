package main

import "fmt"


// The interface that specifies zero methods is known as the *empty* interface
//    interface{}
// An empty interface may hold values of any type. (Every type implements at least zero methods).

// Empty interfaces are used to by code that handles values of unknown type.
// For example, fmt.Print takes any number of arguments of type interface{}

func main() {
	var i interface{}
	describe4(i) // (<nil>, <nil>)

	i = 42
	describe4(i) // (42, int)

	i = "hello"
	describe4(i) // (hello, string)

}

func describe4(i interface{}){
	fmt.Printf("(%v, %T)\n", i, i)
}