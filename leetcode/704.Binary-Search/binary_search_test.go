package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	a := assert.New(t)
	a.Equal(4, search([]int{-1, 0, 3, 5, 9, 12}, 9))
	a.Equal(-1, search([]int{-1, 0, 3, 5, 9, 12}, 2))
}
