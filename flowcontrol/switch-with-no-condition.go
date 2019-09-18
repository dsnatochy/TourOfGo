package main

import (
	"time"
	"fmt"
)

func main() {
	t:=time.Now()
	switch{               // equivalent of "switch true"
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
