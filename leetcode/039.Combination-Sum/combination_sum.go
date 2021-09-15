package combinationsum

func recursiveCombinations(candidates []int, target int, res *[][]int, cur []int) {
	if target == 0 {
		*res = append(*res, append([]int{}, cur...))
		return
	}

	for i, v := range candidates {
		if v <= target {
			recursiveCombinations(candidates[i:], target-v, res, append(cur, candidates[i]))
		}
	}
}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	recursiveCombinations(candidates, target, &res, nil)
	return res
}
