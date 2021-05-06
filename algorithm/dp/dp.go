package dp

var memo []int

// Fibonacci return a fibonacci number
func Fibonacci(n int, bottom bool) int {
	if bottom {
		return bottomup(n)
	}

	memo = make([]int, n+1)
	return topdown(n)
}

func bottomup(n int) int {
	if n <= 1 {
		return n
	}

	first := 1
	second := 1
	var ret int
	for i := 0; i < n-1; i++ {
		ret = first + second
		first = second
		second = ret
	}

	return ret
}

func topdown(n int) int {
	if memo[n] > 0 {
		return memo[n]
	}

	if n <= 1 {
		memo[n] = 1
		return memo[n]
	}

	memo[n] = topdown(n-1) + topdown(n-2)
	return memo[n]
}
