package leetcode

func minDifference(nums []int, queries [][]int) []int {
	N := 101
	dp := make([][]int, len(nums)+1)
	dp[0] = make([]int, N)

	for i, v := range nums {
		dp[i+1] = make([]int, N)
		copy(dp[i+1], dp[i])
		dp[i+1][v]++
	}

	var res []int
	for _, v := range queries {
		var subarray []int
		for j, cnt := range dp[v[0]] {
			if cnt != dp[v[1]+1][j] {
				subarray = append(subarray, j)
			}
		}

		min := -1
		for j := 0; j < len(subarray)-1; j++ {
			if min == -1 || subarray[j+1]-subarray[j] < min {
				min = subarray[j+1] - subarray[j]
			}
		}
		res = append(res, min)
	}

	return res
}
