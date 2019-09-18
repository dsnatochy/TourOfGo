package main

import (
	"time"
	"fmt"
)

// The "error" type is a built-in interface simiar to fmt.Stringer
// type error interface {
//   Error() string
// }

// As with fmt.Stringer, the fmt package looks for the *error* interface when printing values

// Functions often return an *error* value, and calling code should handle errors by testing whether the error equals *nil*

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string{
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

//func (e *MyError) String() string{
//	return "Blah"
//}

func run() error{
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err !=nil {
		fmt.Println(err)  // at 2019-09-17 15:33:17.45624 -0700 PDT m=+0.000253722, it didn't work
	}
}