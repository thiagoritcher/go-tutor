package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (t *rot13Reader) Read(b []byte) (n int, err error){
	for i, v := range b {
		if v > 13 {
			b[i] -= 13
		} else {
			b[i] += 13
		}
	}
	
	if cap(b) == 0 {
		return 0, io.EOF
		
	} else {
		return len(b), nil
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

