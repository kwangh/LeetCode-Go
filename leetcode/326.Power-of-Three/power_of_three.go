package leetcode

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}

func isPowerOfThreeMath(n int) bool {
	return n > 0 && 1162261467%n == 0
}
