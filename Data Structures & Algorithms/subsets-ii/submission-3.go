import "slices"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)

	var dfs func(prefix []int, idx int)
	dfs = func(prefix []int, idx int) {
		result = append(result, slices.Clone(prefix))
		for i := idx; i < len(nums); i++ {
			if i > idx && nums[i] == nums[i-1] {
				continue
			}
			dfs(append(prefix, nums[i]), i+1)
		}
	}
	dfs([]int{}, 0)

	return result
}
