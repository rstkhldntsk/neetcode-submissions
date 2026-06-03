func exist(board [][]byte, word string) bool {
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if k == len(word) {
			return true
		}
		if i >= len(board) || i < 0 || j >= len(board[i]) || j < 0 || board[i][j] != word[k] {
			return false
		}
		tmp := board[i][j]
		board[i][j] = '#'
		res := dfs(i+1, j, k+1) ||
			dfs(i-1, j, k+1) ||
			dfs(i, j+1, k+1) ||
			dfs(i, j-1, k+1)
		board[i][j] = tmp
		return res
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				tmp := board[i][j]
				board[i][j] = '#'
				if dfs(i-1, j, 1) || dfs(i+1, j, 1) || dfs(i, j-1, 1) || dfs(i, j+1, 1) {
					return true
				}
				board[i][j] = tmp
			}
		}
	}
	return false
}
