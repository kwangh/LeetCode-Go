package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsHappy(t *testing.T) {
	a := assert.New(t)
	a.Equal(true, isHappy(19))
	a.Equal(false, isHappy(2))
}
