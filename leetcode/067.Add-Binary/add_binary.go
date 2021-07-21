package leetcode

func addBinary(a string, b string) string {
	var carry byte
	if len(a) < len(b) {
		a, b = b, a
	}
	res := make([]byte, len(a)+1)
	i, j := len(a)-1, len(b)-1
	for j >= 0 || i >= 0 {
		var a1, b1 byte
		if i >= 0 {
			a1 = a[i] - '0'
		}
		if j >= 0 {
			b1 = b[j] - '0'
		}
		add := a1 ^ b1 ^ carry
		carry = (a1^b1)&carry | (a1 & b1)
		res[i+1] = '0' + add
		j--
		i--
	}
	if carry == 1 {
		res[0] = '1'
		return string(res)
	}
	return string(res[1:])
}
