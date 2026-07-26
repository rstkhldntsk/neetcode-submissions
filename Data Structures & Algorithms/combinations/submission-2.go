func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	var backtrack func(idx int, prefix []int)
	backtrack = func(idx int, prefix []int) {
		if len(prefix) == k {
			c := make([]int, k)
			copy(c, prefix)
			result = append(result, c)
			return
		}
		for i := idx; i < n; i++ {
			backtrack(i+1, append(prefix, i+1))
		}
	}
	backtrack(0, make([]int, 0))
	return result
}
