package leetcode

import "testing"

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestMerge(t *testing.T) {
	cases := []struct {
		nums1, nums2, want []int
		m, n               int
	}{
		{nums1: []int{1, 2, 3, 0, 0, 0}, m: 3, nums2: []int{2, 5, 6}, n: 3, want: []int{1, 2, 2, 3, 5, 6}},
		{nums1: []int{1}, m: 1, nums2: []int{}, n: 0, want: []int{1}},
		{nums1: []int{2, 2, 3, 0, 0}, m: 3, nums2: []int{1, 1}, n: 2, want: []int{1, 1, 2, 2, 3}},
	}
	for _, c := range cases {
		merge(c.nums1, c.m, c.nums2, c.n)
		if !equal(c.nums1, c.want) {
			t.Errorf("want:%v instead got:%v", c.want, c.nums1)
		}
	}

}
