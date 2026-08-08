func minPathSum(grid [][]int) int {
	dp := make([][]int, 2)
	for i := range len(dp) {
		dp[i] = make([]int, len(grid[0])+1)
		for j := range len(dp[i]) {
			dp[i][j] = math.MaxInt
		}
	}
	dp[1][0] = 0
	
	for i := range len(grid) {
		for j := range len(grid[0]) {
			dp[1][j+1] = min(dp[0][j+1], dp[1][j]) + grid[i][j]
		}
		dp[0], dp[1] = dp[1], make([]int, len(grid[0])+1)
		dp[1][0] = math.MaxInt
	}
	return dp[0][len(dp[0])-1]
}
