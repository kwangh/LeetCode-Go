package leetcode

const (
	MaxInt32 = 1<<31 - 1
	MinInt32 = -1 << 31
)

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if dividend == MinInt32 && divisor == -1 {
		return MaxInt32
	}

	var negative bool
	if dividend > 0 && divisor < 0 {
		negative = true
		divisor = -divisor
	} else if dividend < 0 && divisor > 0 {
		negative = true
		dividend = -dividend
	} else if dividend < 0 && divisor < 0 {
		dividend = -dividend
		divisor = -divisor
	}

	result := 0
	for dividend >= divisor {
		d, m := divisor, 1
		for d<<1 <= dividend {
			d = d << 1
			m = m << 1
		}
		dividend -= d
		result += m
	}

	if negative {
		result = -result
	}
	return result
}
