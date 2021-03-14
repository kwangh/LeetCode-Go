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
	for sum := 0; sum <= dividend; sum += divisor {
		result++
	}
	result -= 1

	if negative {
		result = -result
	}
	return result
}
