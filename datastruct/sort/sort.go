package sort

import "math/rand"

// SelectionSort returns sorted int slice
// time complexity: O(n^2)
// space complexity: O(n)
func SelectionSort(s []int) {
	for i := 0; i < len(s)-1; i++ {
		min := i
		for j := i + 1; j < len(s); j++ {
			if s[min] > s[j] {
				min = j
			}
		}
		s[min], s[i] = s[i], s[min]
	}
}

// InsertionSort returns sorted int slice
// time complexity: O(n^2)
// space complexity: O(n)
func InsertionSort(s []int) {
	for i := 1; i < len(s); i++ {
		key, j := s[i], i-1
		for j >= 0 && key < s[j] {
			s[j+1] = s[j]
			j--
		}
		s[j+1] = key
	}
}

// BubbleSort returns sorted int slice
// time complexity: O(n^2)
// space complexity: O(n)
func BubbleSort(s []int) {
	for i := 0; i < len(s)-1; i++ {
		for j := 1; j < len(s)-i; j++ {
			if s[j-1] > s[j] {
				s[j-1], s[j] = s[j], s[j-1]
			}
		}
	}
}

func merge(s []int, start int, end int, middle int) {
	tmp := []int{}
	i, j := start, middle+1
	for i <= middle && j <= end {
		if s[i] <= s[j] {
			tmp = append(tmp, s[i])
			i++
		} else {
			tmp = append(tmp, s[j])
			j++
		}
	}

	for i <= middle {
		tmp = append(tmp, s[i])
		i++
	}
	for j <= end {
		tmp = append(tmp, s[j])
		j++
	}

	k := 0
	for i := start; i <= end; i++ {
		s[i] = tmp[k]
		k++
	}
}

// MergeSort returns sorted int slice
// time complexity: O(nlogn)
// space complexity: O(n) = 2n
func MergeSort(s []int, start int, end int) {
	if start < end {
		m := (start + end) / 2
		MergeSort(s, start, m)
		MergeSort(s, m+1, end)
		merge(s, start, end, m)
	}
}

// Pivot returns pivot number which is
// medain of 3 random numbers
func Pivot(s []int) int {
	n := len(s)
	return Median(s[rand.Intn(n)],
		s[rand.Intn(n)],
		s[rand.Intn(n)])
}

// Median returns median of 3 numbers
func Median(a, b, c int) int {
	if a < b {
		switch {
		case b < c:
			return b
		case a < c:
			return c
		default:
			return a
		}
	}
	switch {
	case a < c:
		return a
	case b < c:
		return c
	default:
		return b
	}
}

// Partition reorders the elements of s so that:
// - all elements in s[:low] are less than p,
// - all elements in s[low:high] are equal to p,
// - all elements in s[high:] are greater than p.
func Partition(s []int, p int) (low, high int) {
	low, mid, high := 0, 0, len(s)
	for mid < high {
		// Invariant:
		//  - v[:low] < p
		//  - v[low:mid] = p
		//  - v[mid:high] are unknown
		//  - v[high:] > p
		//
		//         < p       = p        unknown       > p
		//     -----------------------------------------------
		// v: |          |          |a            |           |
		//     -----------------------------------------------
		//                ^          ^             ^
		//               low        mid           high
		switch a := s[mid]; {
		case a < p:
			s[mid], s[low] = s[low], s[mid]
			low++
			mid++
		case a == p:
			mid++
		default: // a > p
			s[mid], s[high-1] = s[high-1], s[mid]
			high--
		}
	}
	return
}

// QuickSort returns sorted int slice
// time complexity: O(nlogn) worst case O(n^2)
// space complexity: O(n)
// https://yourbasic.org/golang/quicksort-optimizations/
func QuickSort(s []int) {
	if len(s) < 20 {
		InsertionSort(s)
		return
	}
	p := Pivot(s)
	low, high := Partition(s, p)
	QuickSort(s[:low])
	QuickSort(s[high:])
}
