package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTitleToNumber(t *testing.T) {
	a := assert.New(t)
	a.Equal(1, titleToNumber("A"), "titleToNumber(\"A\") should be 1")
	a.Equal(28, titleToNumber("AB"), "titleToNumber(\"AB\") should be 28")
	a.Equal(701, titleToNumber("ZY"), "titleToNumber(\"ZY\") should be 701")
	a.Equal(2147483647, titleToNumber("FXSHRXW"), "titleToNumber(\"FXSHRXW\") should be 2147483647")
}
