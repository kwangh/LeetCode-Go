package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestConsecutive(t *testing.T) {
	a := assert.New(t)
	a.Equal(4, longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	a.Equal(9, longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
