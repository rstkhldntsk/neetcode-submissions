func searchMatrix(matrix [][]int, target int) bool {
	l, r := -1, len(matrix)*len(matrix[0])
	for r-l > 1 {
		m := l + (r-l)/2
		row := m / len(matrix[0])
		col := m % len(matrix[0])
		if matrix[row][col] > target {
			r = m
		} else if matrix[row][col] < target {
			l = m
		} else {
			return true
		}
	}
	return false
}
