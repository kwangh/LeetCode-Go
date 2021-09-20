package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	a := assert.New(t)
	a.Equal([][]int{{1, 5}, {6, 9}}, insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	a.Equal([][]int{{1, 2}, {3, 10}, {12, 16}}, insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
	a.Equal([][]int{{5, 7}}, insert([][]int{}, []int{5, 7}))
}
