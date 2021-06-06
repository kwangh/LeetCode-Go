package leetcode

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	var reverse int
	for x > reverse {
		reverse = reverse*10 + x%10
		x /= 10
	}

	return x == reverse || x == reverse/10
}
