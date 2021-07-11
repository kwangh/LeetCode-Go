package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoinChange(t *testing.T) {
	a := assert.New(t)
	a.Equal(3, coinChange([]int{1, 2, 5}, 11))
	a.Equal(-1, coinChange([]int{2}, 3))
	a.Equal(0, coinChange([]int{1}, 0))
	a.Equal(1, coinChange([]int{1}, 1))
	a.Equal(2, coinChange([]int{1}, 2))
}
