package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	a := assert.New(t)
	a.Equal(3, maxProfit([]int{1, 2, 3, 0, 2}))
	a.Equal(0, maxProfit([]int{1}))
}
