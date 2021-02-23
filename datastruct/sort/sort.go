package sort

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
