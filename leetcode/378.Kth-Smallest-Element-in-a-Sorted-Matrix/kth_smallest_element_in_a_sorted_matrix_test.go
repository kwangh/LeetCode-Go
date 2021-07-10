package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthSmallest(t *testing.T) {
	a := assert.New(t)
	a.Equal(13, kthSmallest([][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8))
	a.Equal(-5, kthSmallest([][]int{{-5}}, 1))
}
