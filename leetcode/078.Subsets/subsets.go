package leetcode

func subset(res *[][]int, cur, nums []int) {
	if len(nums) == 0 {
		*res = append(*res, append([]int{}, cur...))
		return
	}
	subset(res, cur, nums[1:])
	subset(res, append(cur, nums[0]), nums[1:])
}

func subsets(nums []int) [][]int {
	var res [][]int
	subset(&res, []int{}, nums)
	return res
}
