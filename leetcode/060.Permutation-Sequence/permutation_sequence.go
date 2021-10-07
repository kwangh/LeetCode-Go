package leetcode

func getPermutation(n int, k int) string {
	seq := make([]byte, n)
	fac := 1
	for i := 1; i <= n; i++ {
		seq[i-1] = byte(i + '0')
		fac *= i
	}
	k--
	res := make([]byte, 0, n)
	for n > 0 {
		fac /= n
		n--
		i := k / fac
		res = append(res, seq[i])
		copy(seq[i:], seq[i+1:])
		k -= fac * i
	}
	return string(res)
}
