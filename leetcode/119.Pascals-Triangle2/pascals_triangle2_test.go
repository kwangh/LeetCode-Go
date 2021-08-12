package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRow(t *testing.T) {
	a := assert.New(t)
	a.Equal([]int{1, 3, 3, 1}, getRow(3))
	a.Equal([]int{1}, getRow(0))
	a.Equal([]int{1, 1}, getRow(1))
}
