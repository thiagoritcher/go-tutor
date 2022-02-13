package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fa, fb := 0, 0;
	return func() int {
		t := fb
		fb = fa + fb
		fa = t
		if fa == 0 && fb == 0 {
			fa = 1
		}
		return fb
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

