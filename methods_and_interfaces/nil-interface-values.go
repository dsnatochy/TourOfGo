package main

import "fmt"

// A nil interface value holds neither value nor concrete type

// Calling a method on a nil interface is a run-time error because there is
// no type inside the interface tuple to indicate which concrete method to call

type I4 interface {
	M4()
}

func main() {
	var i I4
	describe3(i) // (<nil>, <nil>)
	i.M4() // panic: runtime error: invalid memory address or nil pointer dereference
}

func describe3(i I4){
	fmt.Printf("(%v, %T)\n", i, i)
}