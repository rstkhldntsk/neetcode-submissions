func numSquares(n int) int {
	memo := make(map[int]int)
	var dfs func(target int) int
	dfs = func(target int) int {
		if target == 0 {
			return 0
		}
		if val, ok := memo[target]; ok {
			return val
		}
		res := target
		for i := 1; i*i <= target; i++ {
			res = min(res, 1+dfs(target-i*i))
		}
		memo[target] = res
		return res
	}
	return dfs(n)
}
