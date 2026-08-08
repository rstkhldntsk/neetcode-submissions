func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid)+1)
	for i := range len(dp) {
		dp[i] = make([]int, len(grid[0])+1)
		for j := range len(dp[i]) {
			dp[i][j] = math.MaxInt
		}
	}
	dp[1][0] = 0
	
	for i := range len(grid) {
		for j := range len(grid[0]) {
			dp[i+1][j+1] = min(dp[i+1][j], dp[i][j+1]) + grid[i][j]
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}
