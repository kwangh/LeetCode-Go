package leetcode

import "testing"

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

func TestPartition(t *testing.T) {
	cases := []struct {
		s    string
		want [][]string
	}{
		{"aab", [][]string{{"a", "a", "b"}, {"aa", "b"}}},
		{"a", [][]string{{"a"}}},
	}
	for _, c := range cases {
		got := partition(c.s)
		if len(c.want) != len(got) {
			t.Errorf("want:%v instead got:%v", c.want, got)
			break
		}
		for i, w := range c.want {
			if !equal(w, got[i]) {
				t.Errorf("want:%v instead got:%v", c.want, got)
				break
			}
		}
	}
}
