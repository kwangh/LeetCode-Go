package leetcode

import (
	"testing"
)

func TestCountAndSay(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{
			in:   1,
			want: "1",
		},
		{
			in:   2,
			want: "11",
		},
		{
			in:   4,
			want: "1211",
		},
	}
	for _, c := range cases {
		got := countAndSay(c.in)
		if c.want != got {
			t.Errorf("want:%s, instead got:%s\n", c.want, got)
		}
	}
}
