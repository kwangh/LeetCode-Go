package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	a := assert.New(t)
	a.Equal(7, maxProfit([]int{7, 1, 5, 3, 6, 4}), "maxProfit([]int{7,1,5,3,6,4}) should be 7")
	a.Equal(4, maxProfit([]int{1, 2, 3, 4, 5}), "maxProfit([]int{1,2,3,4,5}) should be 4")
	a.Equal(0, maxProfit([]int{7, 6, 4, 3, 1}), "maxProfit([]int{7,6,4,3,1}) should be 0")
	a.Equal(2, maxProfit([]int{2, 4, 1}), "maxProfit([]int{2, 4, 1}) should be 2")
	a.Equal(15, maxProfit([]int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}), "maxProfit([]int{1,2,4,2,5,7,2,4,9,0}) should be 15")
}
