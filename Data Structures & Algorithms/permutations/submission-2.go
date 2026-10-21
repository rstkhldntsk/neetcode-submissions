import "slices"

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	seen := make(map[int]bool)
	
	var dfs func(prefix []int)
	dfs = func(prefix []int) {
		if len(prefix) == len(nums) {
			result = append(result, slices.Clone(prefix))
			return
		}

		for i := 0; i < len(nums); i++ {
			if !seen[nums[i]] {
				seen[nums[i]] = true
				dfs(append(prefix, nums[i]))
				seen[nums[i]] = false
			}
		}
	}
	dfs([]int{})

	return result
}
