package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSumClosest(t *testing.T) {
	a := assert.New(t)
	a.Equal(2, threeSumClosest([]int{-1, 2, 1, -4}, 1))
	a.Equal(0, threeSumClosest([]int{0, 0, 0}, 1))
}
