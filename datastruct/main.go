package main

import (
	"fmt"

	"github.com/kwangh/go-tutorial/datastruct/sort"
)

func main() {
	s1 := []int{2, 5, 2, 1, 4}
	sort.SelectionSort(s1)
	fmt.Println(s1)
	s2 := []int{2, 5, 2, 1, 4}
	sort.InsertionSort(s2)
	fmt.Println(s2)
}
