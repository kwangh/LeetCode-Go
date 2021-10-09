package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindWords(t *testing.T) {
	a := assert.New(t)
	a.Equal([]string{"oath", "eat"}, findWords([][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}, []string{"oath", "pea", "eat", "rain"}))
	a.Equal([]string{}, findWords([][]byte{{'a', 'b'}, {'c', 'd'}}, []string{"abcb"}))
}
