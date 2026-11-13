func exist(board [][]byte, word string) bool {
	var backtrack func(idx int, i, j int) bool
	backtrack = func(idx int, i, j int) bool {
		if idx == len(word) {
			return true
		}
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[idx] {
			return false
		}
		tmp := board[i][j]
		board[i][j] = '#'
		res :=  backtrack(idx+1, i+1, j) ||
			backtrack(idx+1, i-1, j) ||
			backtrack(idx+1, i, j+1) ||
			backtrack(idx+1, i, j-1)
		board[i][j] = tmp
		return res
	}
	
	for i := range board {
		for j := range board[0] {
			if backtrack(0, i, j) {
				return true
			}
		}
	}
	
	return false
}
