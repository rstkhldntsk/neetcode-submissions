import "slices"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := make([][]int, 0)

	var dfs func(prefix []int, idx int, runningSum int)
	dfs = func(prefix []int, idx int, runningSum int) {
		if runningSum == target {
			result = append(result, slices.Clone(prefix))
			return
		}
		for i := idx; i < len(candidates); i++ {
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}
			if runningSum+candidates[i] > target {
				return
			}
			dfs(append(prefix, candidates[i]), i+1, runningSum+candidates[i])
		}
	}
	dfs([]int{}, 0, 0)

	return result
}
