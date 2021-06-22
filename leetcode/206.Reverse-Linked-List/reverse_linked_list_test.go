package leetcode

import (
	"testing"

	"github.com/kwangh/leetcode/structures"
	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	a := assert.New(t)
	a.Equal(structures.Lists([]int{5, 4, 3, 2, 1}), reverseList((structures.Lists([]int{1, 2, 3, 4, 5}))))
	a.Equal(structures.Lists([]int{2, 1}), reverseList((structures.Lists([]int{1, 2}))))
	a.Equal(structures.Lists([]int{}), reverseList((structures.Lists([]int{}))))
}
