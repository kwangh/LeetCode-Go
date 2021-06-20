package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func numberOfRounds(startTime string, finishTime string) int {
	s := int(startTime[0]-'0')*600 + int(startTime[1]-'0')*60 + int(startTime[3]-'0')*10 + int(startTime[4]-'0')
	f := int(finishTime[0]-'0')*600 + int(finishTime[1]-'0')*60 + int(finishTime[3]-'0')*10 + int(finishTime[4]-'0')
	if f < s {
		f += 60 * 24
	}
	return max(0, f/15-(s+14)/15)
}
