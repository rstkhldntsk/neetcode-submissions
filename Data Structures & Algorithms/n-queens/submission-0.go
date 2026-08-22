func solveNQueens(n int) [][]string {
		board := make([][]string, n)
	for i := range n {
		board[i] = make([]string, n)
		for j := range board[i] {
			board[i][j] = "."
		}
	}
	
	result := make([][]string, 0)
	cols := make([]bool, n)
	diag1 := make([]bool, 2*n-1)
	diag2 := make([]bool, 2*n-1)
	var backtrack func(r int)
	backtrack = func(r int) {
		if r == n {
			c := make([]string, 0, n)
			for i := range board {
				c = append(c, strings.Join(board[i], ""))
			}
			result = append(result, c)
			return
		}
		for c := range n {
			d1 := r + c
			d2 := r - c + n - 1
			if !cols[c] && !diag1[d1] && !diag2[d2] {
				cols[c] = true
				diag1[d1] = true
				diag2[d2] = true
				board[r][c] = "Q"
				backtrack(r + 1)
				cols[c] = false
				diag1[d1] = false
				diag2[d2] = false
				board[r][c] = "."
			}
		}
	}
	backtrack(0)
	return result
}