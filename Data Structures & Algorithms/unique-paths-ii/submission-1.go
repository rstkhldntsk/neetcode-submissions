func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([][]int, 2)
	for i := range len(dp) {
		dp[i] = make([]int, len(obstacleGrid[0])+1)
	}
	
	dp[1][0] = 1
	for r := range len(obstacleGrid) {
		for c := range len(obstacleGrid[r]) {
			if obstacleGrid[r][c] != 1 {
				dp[1][c+1] = dp[1][c] + dp[0][c+1]
			}
		}
		dp[0], dp[1] = dp[1], make([]int, len(obstacleGrid[0])+1)
	}
	return dp[0][len(dp[0])-1]
}
