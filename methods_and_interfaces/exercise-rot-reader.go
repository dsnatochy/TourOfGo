package main

import (
	"io"
	"strings"
	"os"
)

type rot13Reader struct{
	r io.Reader
}

// A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.
//
//For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and
// returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).

// Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the
// rot13 substitution cipher to all alphabetical characters.
// https://en.wikipedia.org/wiki/ROT13
func (r rot13Reader) Read(b []byte)(n int, err error){

	n, err = r.r.Read(b)

	if err == nil{
		for i:= 0; i < n; i++ {
			b[i] = rot13(b[i])
		}
	}
	return

}

func rot13(b byte) byte{
	if (b <= 'm' && b >= 'a') || (b <= 'M' && b > 'A') {
		b += 13
	} else if (b > 'M' && b <= 'Z') || (b > 'm' && b <= 'z') {
		b = b - 13
	}
	return b
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

}