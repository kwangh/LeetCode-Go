package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert(t *testing.T) {
	a := assert.New(t)
	a.Equal("PAHNAPLSIIGYIR", convert("PAYPALISHIRING", 3))
	a.Equal("PINALSIGYAHRPI", convert("PAYPALISHIRING", 4))
	a.Equal("A", convert("A", 1))
	a.Equal("AB", convert("AB", 1))
}
