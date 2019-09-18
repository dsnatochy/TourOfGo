package main

import (
	"os"
	"io"
	"fmt"
)

func copy() (written int64, err error) {
	src, err := os.Open("/tmp/source")
	defer src.Close()
	if err != nil{
		return
	}

	dst, err := os.Create("/tmp/dest")
	defer dst.Close()
	if err != nil{
		return
	}


	written, err = io.Copy(dst, src)
	return
}

func main() {
	_, err := copy()
	if err != nil{
		fmt.Println("Copy failed", err)
	}

}