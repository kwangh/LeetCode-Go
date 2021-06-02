package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]int{0, 1}, TwoSum([]int{2, 7, 11, 15}, 9), "TwoSum([2,7,11,15],9) should be [0,1]")
	assert.Equal([]int{1, 2}, TwoSum([]int{3, 2, 4}, 6), "TwoSum([3,2,4],6) should be [1,2]")
	assert.Equal([]int{0, 1}, TwoSum([]int{3, 3}, 6), "TwoSum([3,3],6) should be [0,1]")
}
