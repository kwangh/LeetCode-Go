package main

import (
	"fmt"

	"github.com/kwangh/go-tutorial/datastruct/sort"
)

func main() {
	s := []int{1, 5, 4, 3, 2}
	sort.SelectionSort(s)
	fmt.Println(s)
}
