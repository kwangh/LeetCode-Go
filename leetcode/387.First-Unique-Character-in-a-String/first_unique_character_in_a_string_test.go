package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstUniqChar(t *testing.T) {
	a := assert.New(t)
	a.Equal(0, firstUniqChar("leetcode"))
	a.Equal(2, firstUniqChar("loveleetcode"))
	a.Equal(-1, firstUniqChar("aabb"))
}
