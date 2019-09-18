package main

import "fmt"

//type MyReader interface {
//	Read(b []byte) (int, error)
//}

type MyReader struct{}

// TODO: Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
func (m MyReader) Read(b []byte)(n int, err error){
	for x := range b{
		fmt.Println(x)
		b[x]='A'
	}
	return len(b), nil
}

func main()  {
	a := make([]byte, 10)
	var reader MyReader
	n, _ := reader.Read(a)
	println(n)
}