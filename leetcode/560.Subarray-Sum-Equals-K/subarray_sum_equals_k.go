package leetcode

func subarraySum(nums []int, k int) int {
	var res, sum int
	m := make(map[int]int, len(nums)+1)
	m[0] = 1
	for _, num := range nums {
		sum += num
		res += m[sum-k]
		m[sum]++
	}
	return res
}
