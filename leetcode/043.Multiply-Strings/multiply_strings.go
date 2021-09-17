package leetcode

import "strings"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	nums := make([]int, len(num1)+len(num2))
	for i, digit1 := range num1 {
		for j, digit2 := range num2 {
			nums[i+j+1] += int(digit1-'0') * int(digit2-'0')
		}
	}

	for i := len(nums) - 1; i > 0; i-- {
		nums[i-1] += nums[i] / 10
		nums[i] = nums[i] % 10
	}
	var s strings.Builder
	if nums[0] != 0 {
		s.WriteRune(rune('0' + nums[0]))
	}
	for i := 1; i < len(nums); i++ {
		s.WriteRune(rune('0' + nums[i]))
	}
	return s.String()
}
