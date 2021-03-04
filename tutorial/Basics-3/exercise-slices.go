package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	v := make([][]uint8, dy)

	for y := range v {
		v[y] = make([]uint8, dx)

		for x := range v[y] {
			v[y][x] = uint8((x * y) ^ x - y)
		}
	}

	return v
}

func main() {
	pic.Show(Pic)
}
