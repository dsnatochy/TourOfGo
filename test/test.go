package main

import (
	"fmt"
	"strconv"
)

func test(num int) string{
	return fmt.Sprint(num)
}
func main() {
	a := test(5)
	fmt.Printf("The type is %T\n", a)


	i, err := strconv.Atoi("42")
	if err != nil{

		fmt.Println(err == nil, i)
	}else{
		fmt.Println("Converted: ", i)
	}
}
