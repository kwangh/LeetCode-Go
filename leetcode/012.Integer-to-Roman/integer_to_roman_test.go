package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToRoman(t *testing.T) {
	a := assert.New(t)
	a.Equal("III", intToRoman(3))
	a.Equal("IV", intToRoman(4))
	a.Equal("IX", intToRoman(9))
	a.Equal("LVIII", intToRoman(58))
	a.Equal("MCMXCIV", intToRoman(1994))
}
