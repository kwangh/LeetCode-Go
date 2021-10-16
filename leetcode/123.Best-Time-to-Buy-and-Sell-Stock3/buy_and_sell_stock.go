package leetcode

func maxProfit(prices []int) int {
	buy1, buy2 := prices[0], prices[0]
	sell1, sell2 := 0, 0
	for i := 1; i < len(prices); i++ {
		buy1 = Min(buy1, prices[i])
		sell1 = Max(sell1, prices[i]-buy1)
		buy2 = Min(buy2, prices[i]-sell1)
		sell2 = Max(sell2, prices[i]-buy2)
	}
	return sell2
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
