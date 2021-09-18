package leetcode

func jump(nums []int) int {
	var jump, cur, next int
	for i, n := range nums[:len(nums)-1] {
		if n+i > next {
			next = n + i
		}
		if i == cur {
			jump++
			cur = next
		}
	}
	return jump
}
