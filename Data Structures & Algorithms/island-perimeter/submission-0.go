func islandPerimeter(grid [][]int) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	
	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] == 0 {
			return 1
		}
		if visited[r][c] {
			return 0
		}
		visited[r][c] = true
		return  dfs(r-1, c) + dfs(r+1, c) + dfs(r, c+1) + dfs(r, c-1)
	}
	
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 1 {
				return dfs(r, c)
			}
		}
	}
	
	return 0
}
