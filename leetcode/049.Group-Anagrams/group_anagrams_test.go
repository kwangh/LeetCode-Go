package leetcode

import (
	"testing"
)

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestGroupAnagrams(t *testing.T) {
	cases := []struct {
		strs []string
		want [][]string
	}{
		{
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			strs: []string{"bdddddddddd", "bbbbbbbbbbc"},
			want: [][]string{{"bbbbbbbbbbc"}, {"bdddddddddd"}},
		},
		{
			strs: []string{"a"},
			want: [][]string{{"a"}},
		},
	}
	for _, c := range cases {
		got := groupAnagrams(c.strs)

		if len(got) != len(c.want) {
			t.Errorf("want:%v instead got:%v", c.want, got)
			break
		}

		t.Logf("want:%v got:%v", c.want, got)
	}
}
