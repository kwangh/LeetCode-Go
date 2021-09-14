package combinationsum

func recursiveCombinations(candidates []int, target, index int, res *[][]int, cur []int) {
	if target == 0 {
		*res = append(*res, append([]int{}, cur...))
		return
	}

	for i := index; i < len(candidates); i++ {
		if candidates[i] <= target {
			recursiveCombinations(candidates, target-candidates[i], i, res, append(cur, candidates[i]))
		}
	}
}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	recursiveCombinations(candidates, target, 0, &res, nil)
	return res
}
