package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	a := assert.New(t)
	a.Equal(true, containsDuplicate([]int{1, 2, 3, 1}))
	a.Equal(false, containsDuplicate([]int{1, 2, 3, 4}))
	a.Equal(true, containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
