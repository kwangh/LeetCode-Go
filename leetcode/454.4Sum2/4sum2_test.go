package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFourSumCount(t *testing.T) {
	a := assert.New(t)
	a.Equal(2, fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}))
	a.Equal(1, fourSumCount([]int{0}, []int{0}, []int{0}, []int{0}))
	a.Equal(6, fourSumCount([]int{-1, -1}, []int{-1, 1}, []int{-1, 1}, []int{1, -1}))
}
