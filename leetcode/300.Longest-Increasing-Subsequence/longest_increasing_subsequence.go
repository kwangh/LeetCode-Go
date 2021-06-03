package leetcode

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	len := 0
	for _, v := range nums {
		i, j := 0, len
		for i < j {
			m := (i + j) / 2
			if dp[m] < v {
				i = m + 1
			} else {
				j = m
			}
		}
		dp[i] = v
		if i == len {
			len++
		}
	}

	return len
}
