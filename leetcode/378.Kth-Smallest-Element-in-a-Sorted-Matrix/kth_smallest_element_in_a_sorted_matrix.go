package leetcode

func kthSmallest(matrix [][]int, k int) int {
	lo, hi := matrix[0][0], matrix[len(matrix)-1][len(matrix[0])-1]
	for lo < hi {
		mid, cnt, j := (lo+hi)/2, 0, len(matrix[0])-1
		for i := 0; i < len(matrix); i++ {
			for j >= 0 && matrix[i][j] > mid {
				j--
			}
			cnt += j + 1
		}
		if cnt < k {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
