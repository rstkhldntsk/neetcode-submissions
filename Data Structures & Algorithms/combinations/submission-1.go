import "slices"

func combine(n int, k int) [][]int {
    combinations := make([][]int, 0)
	var dfs func(comb []int, i int)
	dfs = func(comb []int, i int) {
		if len(comb) == k {
			combinations = append(combinations, slices.Clone(comb))
			return
		}
		for j := i; j <= n; j++ {
			dfs(append(comb, j), j+1)
		}
	}
	dfs([]int{}, 1)
	return combinations
}