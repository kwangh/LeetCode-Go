package leetcode

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{
			[]int{1, 1, 2},
			2,
		},
		{
			[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			5,
		},
		{
			[]int{},
			0,
		},
	}

	for _, c := range cases {
		got := removeDuplicates(c.in)
		t.Logf("c.in: %v", c.in)
		if got != c.want {
			t.Errorf("want: %d, instead got: %d, c.in: %v", c.want, got, c.in)
		}
	}
}
