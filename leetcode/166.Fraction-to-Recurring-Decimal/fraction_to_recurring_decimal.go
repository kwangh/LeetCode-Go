package leetcode

import (
	"strconv"
)

func contains(a []int, b int) bool {
	if len(a) == 0 {
		return false
	}

	for _, v := range a {
		if v == b {
			return true
		}
	}
	return false
}

func fractionToDecimal(numerator int, denominator int) string {
	var res string
	if numerator < 0 && denominator > 0 {
		numerator = -numerator
		res = "-"
	} else if numerator > 0 && denominator < 0 {
		denominator = -denominator
		res = "-"
	} else if numerator < 0 && denominator < 0 {
		numerator = -numerator
		denominator = -denominator
	}
	res += strconv.Itoa(numerator / denominator)
	numerator %= denominator
	if numerator == 0 {
		return res
	}

	res += "."
	var stack []int
	for !contains(stack, numerator) {
		stack = append(stack, numerator)
		numerator = (numerator * 10) % denominator
		if numerator == 0 {
			for _, v := range stack {
				res += strconv.Itoa(v * 10 / denominator)
			}
			return res
		}
	}
	for _, v := range stack {
		if v == numerator {
			res += "("
		}
		res += strconv.Itoa(v * 10 / denominator)
	}
	res += ")"
	return res
}
