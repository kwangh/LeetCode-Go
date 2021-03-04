package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(i, z)
	}
	return z
}

func main() {
	v := 5.0
	fmt.Println(Sqrt(v))
	fmt.Println(math.Sqrt(v))
}
