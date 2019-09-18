package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os:= runtime.GOOS; os{
	case "darwin":                  //  (1) break statement is implicit (2) does not need to be a constant or integer
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd
		// plan9, windows
		fmt.Printf("%s.\n", os)
	}
}
