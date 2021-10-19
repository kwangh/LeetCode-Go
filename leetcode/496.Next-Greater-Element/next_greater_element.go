package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	js := make(map[int]int, len(nums2))
	stack := make([]int, 0, len(nums2))
	for _, num2 := range nums2 {
		for len(stack) != 0 && stack[len(stack)-1] < num2 {
			js[stack[len(stack)-1]] = num2
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num2)
	}
	res := make([]int, len(nums1))
	for i, num1 := range nums1 {
		if v, ok := js[num1]; ok {
			res[i] = v
		} else {
			res[i] = -1
		}
	}
	return res
}
