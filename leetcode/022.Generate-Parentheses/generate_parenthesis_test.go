package leetcode

import (
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	cases := []struct {
		n    int
		want []string
	}{
		{
			n:    3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			n:    1,
			want: []string{"()"},
		},
	}
	for _, c := range cases {
		got := GenerateParenthesis(c.n)
		if len(got) != len(c.want) {
			t.Errorf("input: %d\texpected %v instead got %v\n", c.n, c.want, got)
		}
	}
}
