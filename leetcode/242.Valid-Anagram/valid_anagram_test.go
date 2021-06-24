package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	a := assert.New(t)
	a.Equal(true, isAnagram("anagram", "nagaram"))
	a.Equal(false, isAnagram("rat", "car"))
}
