package leetcode

func countPrimes(n int) int {
	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i < n; i++ {
		if !isPrime[i] {
			continue
		}
		for j := i * i; j < n; j += i {
			isPrime[j] = false
		}

	}
	var cnt int
	for i := 2; i < n; i++ {
		if isPrime[i] {
			cnt++
		}
	}
	return cnt
}
