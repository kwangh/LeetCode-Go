package leetcode

func maxProfit(prices []int) int {
	var stay, sell int
	buy, l := -prices[0], len(prices)-1
	for i := 1; i <= l; i++ {
		tmp := stay
		stay = max(stay, sell)
		sell = buy + prices[i]
		buy = max(buy, tmp-prices[i])
	}
	return max(stay, sell)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
