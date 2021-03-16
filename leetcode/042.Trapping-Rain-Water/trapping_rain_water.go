package leetcode

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

func trap(height []int) int {
	return trapStack(height)
}

func trapStack(height []int) int {
	res, current := 0, 0
	stack := []int{}
	for current < len(height) {
		for len(stack) != 0 && height[current] > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			distance := current - stack[len(stack)-1] - 1
			boundedHeight := min(height[stack[len(stack)-1]], height[current]) - height[top]
			res += distance * boundedHeight
		}

		stack = append(stack, current)
		current++
	}
	return res
}
