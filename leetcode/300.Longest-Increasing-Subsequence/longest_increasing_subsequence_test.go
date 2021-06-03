package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLIS(t *testing.T) {
	a := assert.New(t)
	a.Equal(4, lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}), "lengthOfLIS([]int{10,9,2,5,3,7,101,18}) should be 4")
	a.Equal(4, lengthOfLIS([]int{0, 1, 0, 3, 2, 3}), "lengthOfLIS([]int{0,1,0,3,2,3}) should be 4")
	a.Equal(1, lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}), "lengthOfLIS([]int{7,7,7,7,7,7,7}) should be 1")
	a.Equal(6, lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}), "lengthOfLIS([]int{1,3,6,7,9,4,10,5,6}) should be 6")
}
