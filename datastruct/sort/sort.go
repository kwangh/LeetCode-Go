package sort

// SelectionSort return sorted int slices
// time complexity: O(n^2)
// space complexity: O(n)
func SelectionSort(s []int) {
	for i := 0; i < len(s)-1; i++ {
		min := i
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				min = j
			}
		}
		s[min], s[i] = s[i], s[min]
	}
}
