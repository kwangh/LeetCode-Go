package leetcode

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChange(coins []int, amount int) int {
	m := make([]int, amount+1)
	for i := range m {
		m[i] = amount + 1
	}
	m[0] = 0
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			m[i] = min(m[i], m[i-coin]+1)
		}
	}
	if m[amount] == amount+1 {
		return -1
	}
	return m[amount]
}
