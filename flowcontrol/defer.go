package main

import "fmt"

func main() {
	// The deferred call's arguments are evaluated immediately,
	// but the function call is not executed until the surrounding function returns.
	defer fmt.Println("This will print last")
	defer fmt.Println("This will print before last")


	fmt.Println("Hello")
}
