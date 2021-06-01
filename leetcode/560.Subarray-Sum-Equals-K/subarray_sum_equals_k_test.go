package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubarraySum(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(2, subarraySum([]int{1, 1, 1}, 2), "subarraySum([1,1,1],2) should be 2")
	assert.Equal(2, subarraySum([]int{1, 2, 3}, 3), "subarraySum([1,2,3],3) should be 2")
}
