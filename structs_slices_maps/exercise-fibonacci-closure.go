package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int

func fibonacci() func() int {
	started := false
	var num, previousNum int
	return func() int{
		if started == false {
			started = true
			num = 1
			previousNum = 0
			return 0;
		}else{
			//fmt.Printf("num: %d, previousNum: %d\n", num, previousNum)
			num, previousNum = num+previousNum, num
			//fmt.Printf("num: %d, previousNum: %d\n", num, previousNum)
			return previousNum
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i<10; i++{
		fmt.Println(f())
	}
}