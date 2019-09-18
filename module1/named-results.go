package module1

import "fmt"

func split(sum int)(x, y int){ // "named" return values - equates to declaring them at the top of func
	x = sum * 4 /9
	y = sum - x
	return  // "naked" return returns named return values
}

func main() {
	fmt.Println(split(17))
}