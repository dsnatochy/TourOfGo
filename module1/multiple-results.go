package module1

import "fmt"

func swap(x, y string) (string, string){// function can return many results.
	return y, x
}

func main() {
	a, b := swap("first", "second")  // := - initialization
	fmt.Println(a, b)

}