import "slices"

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	
	var dfs func(prefix []int, idx int)
	dfs = func(prefix []int, idx int) {
		result = append(result, slices.Clone(prefix))
		for i := idx; i < len(nums); i++ {
			dfs(append(prefix, nums[i]), i+1)
		}
	}
	dfs([]int{}, 0)

	return result
}
