package leetcode

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findDuplicates(nums []int) []int {
	res := []int{}
	for _, v := range nums {
		index := abs(v) - 1
		if nums[index] < 0 {
			res = append(res, index+1)
		}
		nums[index] = -nums[index]
	}
	return res
}
