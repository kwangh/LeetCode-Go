package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	a := assert.New(t)
	a.Equal([][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}, generate(5), "generate(5) should be {1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}")
	a.Equal([][]int{{1}}, generate(1), "generate(1) should be {1}")
}
