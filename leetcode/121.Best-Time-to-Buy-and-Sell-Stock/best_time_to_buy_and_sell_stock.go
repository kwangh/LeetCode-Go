package leetcode

func maxProfit(prices []int) int {
	var res int
	l := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < l {
			l = prices[i]
		} else if prices[i] > l {
			if prices[i]-l > res {
				res = prices[i] - l
			}
		}
	}
	return res
}
