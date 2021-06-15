package leetcode

func sum(a int) int {
	sum := 0
	for a != 0 {
		sum += (a % 10) * (a % 10)
		a /= 10
	}
	return sum
}

func isHappy(n int) bool {
	a, b := n, n
	b = sum(b)
	for a != b {
		a = sum(a)
		b = sum(b)
		b = sum(b)
	}
	return a == 1
}
