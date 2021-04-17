package leetcode

import "testing"

func TestWordBreak(t *testing.T) {
	cases := []struct {
		s        string
		wordDict []string
		want     bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
		{"aaaaaaa", []string{"aaaa", "aaa"}, true},
	}
	for _, c := range cases {
		got := wordBreak(c.s, c.wordDict)
		if c.want != got {
			t.Errorf("string:%s\twant:%t instead got:%t", c.s, c.want, got)
		}
	}
}
