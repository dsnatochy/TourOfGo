package module1

import "fmt"

func add(x int, y int) int{ // return type is after func name or var name
	return x+ y

}

func main() {
	fmt.Println(add(42, 13))
}