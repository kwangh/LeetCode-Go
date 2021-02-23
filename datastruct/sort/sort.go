package sort

// SelectionSort return sorted int slices
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

// InsertionSort return sorted int slices
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
