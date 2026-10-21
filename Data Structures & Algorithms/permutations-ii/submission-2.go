import "slices"

func permuteUnique(nums []int) [][]int {
	counter := map[int]int{}
	for _, num := range nums {
		counter[num]++
	}

	result := make([][]int, 0)
	var dfs func(prefix []int)
	dfs = func(prefix []int) {
		if len(prefix) == len(nums) {
			result = append(result, slices.Clone(prefix))
			return
		}

		for num, cnt := range counter {
			if cnt == 0 {
				continue
			}
			counter[num]--
			dfs(append(prefix, num))
			counter[num]++
		}
	}
	dfs([]int{})

	return result
}
