package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int{
	var words = strings.Fields(s)
	theMap := make(map[string]int)
	for _, v:= range words{
		theMap[v] = theMap[v]+1
	}
	return theMap
}

func main() {
	fmt.Println(WordCount("a a bb bb ccc ccc ccc"))
}
