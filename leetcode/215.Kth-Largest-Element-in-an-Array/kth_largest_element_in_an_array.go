package leetcode

func partition(nums []int, l, r int) int {
	pivot := nums[r]
	i := l
	for j := l; j < r; j++ {
		if nums[j] >= pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[r] = nums[r], nums[i]
	return i
}

func findKthLargest(nums []int, k int) int {
	l, r := 0, len(nums)-1
	k = k - 1
	for l < r {
		p := partition(nums, l, r)
		if p < k {
			l = p + 1
		} else if p > k {
			r = p - 1
		} else {
			break
		}
	}
	return nums[k]
}
