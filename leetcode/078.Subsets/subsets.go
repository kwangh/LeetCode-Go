package leetcode

func subset(res *[][]int, cur, nums []int) {
	if len(nums) == 0 {
		var tmp []int
		*res = append(*res, append(tmp, cur...))
		return
	}
	subset(res, cur, nums[1:])
	subset(res, append(cur, nums[0]), nums[1:])
}

func subsets(nums []int) [][]int {
	var res [][]int
	var cur []int
	subset(&res, cur, nums)
	return res
}
