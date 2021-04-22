package leecode

import "testing"

func TestEvalRPN(t *testing.T) {
	cases := []struct {
		tokens []string
		want   int
	}{
		{[]string{"2", "1", "+", "3", "*"}, 9},
		{[]string{"4", "13", "5", "/", "+"}, 6},
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}
	for _, c := range cases {
		got := evalRPN(c.tokens)
		if c.want != got {
			t.Errorf("tokens:%v\twant:%d instead got:%d", c.tokens, c.want, got)
		}
	}
}
