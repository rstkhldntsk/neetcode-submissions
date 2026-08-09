func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([][]int, len(obstacleGrid)+1)
	for i := range len(dp) {
		dp[i] = make([]int, len(obstacleGrid[0])+1)
	}
	
	dp[1][0] = 1
	for r := range len(obstacleGrid) {
		for c := range len(obstacleGrid[r]) {
			if obstacleGrid[r][c] == 1 {
				continue
			}
			dp[r+1][c+1] = dp[r+1][c] + dp[r][c+1]
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}