package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	a := assert.New(t)
	a.Equal("6", multiply("2", "3"))
	a.Equal("56088", multiply("123", "456"))
	a.Equal("0", multiply("0", "0"))
	a.Equal("419254329864656431168468", multiply("498828660196", "840477629533"))
}
