package main

import (
	"fmt"

	"github.com/kwangh/go-tutorial/datastruct/dp"
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
	}
	fmt.Println("before SelectionSort", slices[0].slice)
	sort.SelectionSort(slices[0].slice)
	fmt.Println("after SelectionSort", slices[0].slice)

	fmt.Println("before InsertionSort", slices[1].slice)
	sort.InsertionSort(slices[1].slice)
	fmt.Println("after InsertionSort", slices[1].slice)

	fmt.Println("before MergeSort", slices[2].slice)
	sort.MergeSort(slices[2].slice, 0, len(slices[2].slice)-1)
	fmt.Println("after MergeSort", slices[2].slice)

	fmt.Println("before QuickSort", slices[3].slice)
	sort.QuickSort(slices[3].slice, 0, len(slices[3].slice)-1)
	fmt.Println("after QuickSort", slices[3].slice)

	fmt.Println("fibonacci bottomup", dp.Fibonacci(28, true))
	fmt.Println("fibonacci topdown", dp.Fibonacci(14, false))
}
