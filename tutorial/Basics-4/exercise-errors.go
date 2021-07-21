package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	z := x / 2
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(i, z)
	}
	return z, nil
}

func main() {
	v, err := Sqrt(2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}

	v, err = Sqrt(-2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}
}
