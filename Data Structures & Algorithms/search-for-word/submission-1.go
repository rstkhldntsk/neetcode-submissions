func exist(board [][]byte, word string) bool {
	b := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		b[i] = make([]bool, len(board[i]))
	}

	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if k == len(word) {
			return true
		}
		if i >= len(board) || i < 0 || j >= len(board[i]) || j < 0 || b[i][j] || board[i][j] != word[k] {
			return false
		}
		b[i][j] = true
		res := dfs(i+1, j, k+1) ||
			dfs(i-1, j, k+1) ||
			dfs(i, j+1, k+1) ||
			dfs(i, j-1, k+1)
		b[i][j] = false
		return res
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				b[i][j] = true
				if dfs(i-1, j, 1) || dfs(i+1, j, 1) || dfs(i, j-1, 1) || dfs(i, j+1, 1) {
					return true
				}
				b[i][j] = false
			}
		}
	}
	return false
}
