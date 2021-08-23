package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertToTitle(t *testing.T) {
	a := assert.New(t)
	a.Equal("A", convertToTitle(1))
	a.Equal("AB", convertToTitle(28))
	a.Equal("ZY", convertToTitle(701))
	a.Equal("FXSHRXW", convertToTitle(2147483647))
}
