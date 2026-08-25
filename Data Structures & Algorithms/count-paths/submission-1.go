func uniquePaths(m int, n int) int {
    grid := make([][]int, m+1)
	for i := range grid {
		grid[i] = make([]int, n+1)
	}
	grid[1][0] = 1
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			grid[i][j] = grid[i][j-1] + grid[i-1][j]
		}	
	}
	return grid[m][n]
}
