package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySqrt(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(2, mySqrt(4), "mySqrt(4) should be 2")
	assert.Equal(2, mySqrt(8), "mySqrt(8) should be 2")
	assert.Equal(1, mySqrt(1), "mySqrt(1) should be 1")
	assert.Equal(1, mySqrt(2), "mySqrt(2) should be 1")
}
