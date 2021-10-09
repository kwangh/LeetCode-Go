package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckInclusion(t *testing.T) {
	a := assert.New(t)
	//a.Equal(true, checkInclusion("ab", "eidbaooo"))
	//a.Equal(false, checkInclusion("ab", "eidboaoo"))
	a.Equal(false, checkInclusion("hello", "ooolleoooleh"))
}
