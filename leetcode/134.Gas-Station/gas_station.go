package leetcode

func canCompleteCircuit(gas []int, cost []int) int {
	var total, tank, start int
	for i, v := range gas {
		tank += v - cost[i]
		if tank < 0 {
			start = i + 1
			total += tank
			tank = 0
		}
	}
	if total+tank < 0 {
		return -1
	}
	return start
}
