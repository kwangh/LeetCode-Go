package main

import (
	"fmt"

	"github.com/kwangh/leetcode/datastruct/dp"
	"github.com/kwangh/leetcode/datastruct/sort"
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
	for _, s := range slices {
		tmp := s
		fmt.Println("Slice before sorting", tmp.slice)
		sort.SelectionSort(tmp.slice)
		fmt.Printf("SelectionSort:\t%v\n", tmp.slice)

		tmp = s
		sort.InsertionSort(tmp.slice)
		fmt.Printf("InsertionSort:\t%v\n", tmp.slice)

		tmp = s
		sort.MergeSort(tmp.slice, 0, len(tmp.slice)-1)
		fmt.Printf("MergeSort:\t%v\n", tmp.slice)

		tmp = s
		sort.QuickSort(tmp.slice)
		fmt.Printf("QuickSort:\t%v\n", tmp.slice)
	}

	fmt.Println("fibonacci bottomup", dp.Fibonacci(28, true))
	fmt.Println("fibonacci topdown", dp.Fibonacci(14, false))
}
