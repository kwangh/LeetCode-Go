package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextPermutation(t *testing.T) {
	a := assert.New(t)
	input := []int{1, 2, 3}
	nextPermutation(input)
	a.Equal([]int{1, 3, 2}, input)

	input = []int{3, 2, 1}
	nextPermutation(input)
	a.Equal([]int{1, 2, 3}, input)

	input = []int{1, 1, 5}
	nextPermutation(input)
	a.Equal([]int{1, 5, 1}, input)

	input = []int{1}
	nextPermutation(input)
	a.Equal([]int{1}, input)

}
