package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfThree(t *testing.T) {
	a := assert.New(t)
	a.Equal(true, isPowerOfThree(27))
	a.Equal(false, isPowerOfThree(0))
	a.Equal(true, isPowerOfThree(9))
	a.Equal(false, isPowerOfThree(45))
}
