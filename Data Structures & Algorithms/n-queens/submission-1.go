func solveNQueens(n int) [][]string {
	board := make([][]string, n)
	for i := range board {
		board[i] = make([]string, n)
		for j := range board[i] {
			board[i][j] = "."
		}
	}
	
	result := make([][]string, 0)
	cols := make([]bool, n)
	posDiag := make([]bool, 2*n) // r+c
	negDiag := make([]bool, 2*n) // r-c
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
			if !cols[c] && !posDiag[r+c] && !negDiag[r-c+n] {
				cols[c] = true
				posDiag[r+c] = true
				negDiag[r-c+n] = true
				board[r][c] = "Q"
				backtrack(r + 1)
				cols[c] = false
				posDiag[r+c] = false
				negDiag[r-c+n] = false
				board[r][c] = "."
			}
		}
	}
	backtrack(0)
	
	return result
}
