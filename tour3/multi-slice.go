package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ddx := make([][]uint8, dx)
	
	for x:= 0; x < dx; x++ {
		ddy := make([]uint8, dy)
		for y:= 0; y < dy; y++ {
			ddy[y] = uint8((x+y)/2)
		}
		ddx[x] = ddy
	}
	return ddx
}

func main() {
	pic.Show(Pic)
}

