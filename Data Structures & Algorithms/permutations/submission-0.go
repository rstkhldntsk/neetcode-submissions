func permute(nums []int) [][]int {
	permutes := make([][]int, 0)
	seen := make(map[int]bool)

	var dfs func(p []int)
	dfs = func(p []int) {
		if len(p) == len(nums) {
			permutes = append(permutes, p)
			return
		}

		for j := 0; j < len(nums); j++ {
			if !seen[nums[j]] {
				seen[nums[j]]=true
				dfs(append(p, nums[j]))
				seen[nums[j]]=false
			}
		}
	}

	dfs([]int{})
	return permutes
}
