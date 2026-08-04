func combinationSum4(nums []int, target int) int {
	memo := make([]int, target+1)
	for i := range memo {
		memo[i] = -1
	}
	
	var dfs func(int) int
	dfs = func(runningSum int) int {
		if runningSum == target {
			return 1
		}
		if runningSum > target {
			return 0
		}
		if memo[runningSum] != -1 {
			return memo[runningSum]
		}
		res := 0
		for i := 0; i < len(nums); i++ {
			res += dfs(runningSum + nums[i])
		}
		memo[runningSum] = res
		return res
	}
	return dfs(0)
}