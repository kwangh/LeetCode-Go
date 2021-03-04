package main

import "fmt"

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		b, a = a, a+b
		return b
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
