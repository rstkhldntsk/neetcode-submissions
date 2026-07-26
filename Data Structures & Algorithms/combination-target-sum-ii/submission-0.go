func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := make([][]int, 0)
	var backtrack func(idx, runningSum int, prefix []int)
	backtrack = func(idx, runningSum int, prefix []int) {
		if runningSum == target {
			c := make([]int, len(prefix))
			copy(c, prefix)
			result = append(result, c)
			return
		}
		for i := idx; i < len(candidates); i++ {
			if runningSum+candidates[i] > target {
				return
			}
            if i > idx && candidates[i] == candidates[i-1] {
                continue
            }
			backtrack(i+1, runningSum+candidates[i], append(prefix, candidates[i]))
		}
	}
	backtrack(0, 0, make([]int, 0))
	return result
}
