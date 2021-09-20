package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumDistinct(t *testing.T) {
	a := assert.New(t)
	a.Equal(3, numDistinct("rabbbit", "rabbit"))
	a.Equal(5, numDistinct("babgbag", "bag"))
}
