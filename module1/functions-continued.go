package module1

import "fmt"

func add2(x, y int) int{ // param type can be ommitted when next param has the same type
	return x+ y

}

func main() {
	fmt.Println(add2(42, 13))
}