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
		{[]int{7, 1, 2, 5, 2}},
		{[]int{9, 5, 2, 3, 1}},
	}
	sort.SelectionSort(slices[0].slice)
	fmt.Println(slices[0].slice)

	sort.InsertionSort(slices[1].slice)
	fmt.Println(slices[1].slice)

	sort.InsertionSort(slices[2].slice)
	fmt.Println(slices[2].slice)

	sort.MergeSort(slices[3].slice, 0, len(slices[3].slice)-1)
	fmt.Println(slices[3].slice)

	sort.QuickSort(slices[4].slice, 0, len(slices[4].slice)-1)
	fmt.Println(slices[4].slice)
}
