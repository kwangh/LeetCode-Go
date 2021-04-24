package leetcode

import "testing"

func TestFractionToDecimal(t *testing.T) {
	cases := []struct {
		numerator, denominator int
		want                   string
	}{
		{numerator: 1, denominator: 2, want: "0.5"},
		{numerator: 2, denominator: 1, want: "2"},
		{numerator: 2, denominator: 3, want: "0.(6)"},
		{numerator: 4, denominator: 333, want: "0.(012)"},
		{numerator: 1, denominator: 333, want: "0.(003)"},
		{numerator: 1, denominator: 5, want: "0.2"},
		{numerator: 1, denominator: 6, want: "0.1(6)"},
		{numerator: -50, denominator: 8, want: "-6.25"},
		{numerator: 7, denominator: -12, want: "-0.58(3)"},
	}
	for _, c := range cases {
		got := fractionToDecimal(c.numerator, c.denominator)
		if c.want != got {
			t.Errorf("%d/%d=\"%s\"(want) instead \"%s\"(got)", c.numerator, c.denominator, c.want, got)
		}
	}
}
