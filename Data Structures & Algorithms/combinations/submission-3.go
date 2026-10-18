func combine(n int, k int) [][]int {
	result := make([][]int, 0)

	var dfs func(prefix []int, idx int)
	dfs = func(prefix []int, idx int) {
		if len(prefix) == k {
			c := make([]int, len(prefix))
			copy(c, prefix)
			result = append(result, c)
			return
		}
		for i := idx; i <= n; i++ {
			dfs(append(prefix, i), i+1)
		}
	}
	dfs([]int{}, 1)

	return result
}
