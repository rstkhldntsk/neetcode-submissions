func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	
	var dfs func(prefix []int, idx int)
	dfs = func(prefix []int, idx int) {
		c := make([]int, len(prefix))
		copy(c, prefix)
		result = append(result, c)
		for i := idx; i < len(nums); i++ {
			dfs(append(prefix, nums[i]), i+1)
		}
	}
	dfs([]int{}, 0)

	return result
}
