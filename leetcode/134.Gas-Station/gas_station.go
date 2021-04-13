package leetcode

func canCompleteCircuit(gas []int, cost []int) int {
	for i, v := range gas {
		if v >= 0 {
			var sum int
			for j := 0; j < len(gas); j++ {
				sum += gas[(i+j)%len(gas)] - cost[(i+j)%len(gas)]
				if sum < 0 {
					break
				}
			}
			if sum >= 0 {
				return i
			}
		}
	}
	return -1
}
