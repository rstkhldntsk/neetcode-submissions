func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	var backtrack func(start int, prefix []int)
	backtrack = func(start int, prefix []int) {
		c := make([]int, len(prefix))
		copy(c, prefix)
		result = append(result, c)
		for i := start; i < len(nums); i++ {
			backtrack(i+1, append(prefix, nums[i]))
		}
	}
	backtrack(0, make([]int, 0))
	return result
}
