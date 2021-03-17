package leetcode

type slope struct {
	xd int
	yd int
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func slopeCoprime(x1, y1, x2, y2 int) slope {
	xd := x1 - x2
	yd := y1 - y2

	if xd == 0 {
		return slope{xd: 0, yd: 1}
	} else if yd == 0 {
		return slope{xd: 1, yd: 0}
	}

	gcd := GCD(xd, yd)
	return slope{xd / gcd, yd / gcd}
}

func maxPoints(points [][]int) int {
	if len(points) < 3 {
		return len(points)
	}

	res := 1
	for i := 0; i < len(points)-2; i++ {
		slopes := make(map[slope]int)
		for j := i + 1; j < len(points); j++ {
			slopes[slopeCoprime(points[i][0], points[i][1], points[j][0], points[j][1])]++
		}

		for _, v := range slopes {
			res = max(res, v)
		}
	}

	return res + 1
}
