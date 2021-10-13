package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombine(t *testing.T) {
	a := assert.New(t)
	a.Equal([][]int{{1}}, combine(1, 1))
}
