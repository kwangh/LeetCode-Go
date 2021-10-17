package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNetworkBecomesIdle(t *testing.T) {
	a := assert.New(t)
	a.Equal(67, networkBecomesIdle([][]int{{5, 7}, {15, 18}, {12, 6}, {5, 1}, {11, 17}, {3, 9}, {6, 11}, {14, 7}, {19, 13}, {13, 3}, {4, 12}, {9, 15}, {2, 10}, {18, 4}, {5, 14}, {17, 5}, {16, 2}, {7, 1}, {0, 16}, {10, 19}, {1, 8}}, []int{0, 2, 1, 1, 1, 2, 2, 2, 2, 1, 1, 1, 2, 1, 1, 1, 1, 2, 1, 1}))
}
