package leetcode

import (
	"math"
	"testing"
)

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func TestMyPow(t *testing.T) {
	cases := []struct {
		x    float64
		n    int
		want float64
	}{
		{2, 10, 1024},
		{2.1, 3, 9.261},
		{2, -2, 0.25},
		{-1, MinInt32, 1},
	}
	for _, c := range cases {
		got := myPow(c.x, c.n)
		if !almostEqual(got, c.want) {
			t.Errorf("want:%f instead got:%f", c.want, got)
		}
	}
}
