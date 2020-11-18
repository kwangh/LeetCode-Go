package leetcode

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	cases := []struct {
		nums1, nums2 []int
		want         float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
	}
	for _, c := range cases {
		got := FindMedianSortedArrays(c.nums1, c.nums2)
		if got != c.want {
			t.Error()
		}
	}
}
