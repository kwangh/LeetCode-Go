package main

import (
	"fmt"

	"github.com/kwangh/go-tutorial/datastruct/sort"
)

func main() {
	slices := []struct {
		slice []int
	}{
		{[]int{2, 5, 2, 1, 4}},
		{[]int{1, 3, 2, 5, 1}},
		{[]int{4, 5, 3, 1, 2}},
	}
	sort.SelectionSort(slices[0].slice)
	fmt.Println(slices[0].slice)

	sort.InsertionSort(slices[1].slice)
	fmt.Println(slices[1].slice)

	sort.InsertionSort(slices[2].slice)
	fmt.Println(slices[2].slice)

	sort.MergeSort(slices[2].slice, 0, len(slices[2].slice)-1)
	fmt.Println(slices[2].slice)
}
