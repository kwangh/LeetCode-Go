package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddBinary(t *testing.T) {
	a := assert.New(t)
	a.Equal("100", addBinary("11", "1"))
	a.Equal("0", addBinary("0", "0"))
	a.Equal("10101", addBinary("1010", "1011"))
}
