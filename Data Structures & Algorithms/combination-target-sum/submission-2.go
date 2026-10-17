func combinationSum(nums []int, target int) [][]int {
	sort.Ints(nums)
    result := make([][]int, 0)

	var dfs func(prefix []int, runningSum int, idx int)
	dfs = func(prefix []int, runningSum int, idx int) {
		if runningSum == target {
			c := make([]int, len(prefix))
			copy(c, prefix)
			result = append(result, c)
			return
		}

		for i := idx; i < len(nums); i++ {
			if runningSum+nums[i] > target {
				return
			}
			dfs(append(prefix, nums[i]), runningSum+nums[i], i)
		}
	}
	dfs(make([]int, 0), 0, 0)

	return result
}
