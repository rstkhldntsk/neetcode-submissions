type NumMatrix struct {
	prefixMatrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	prefix := make([][]int, len(matrix)+1)
	prefix[0] = make([]int, len(matrix[0])+1)
	for i := 0; i < len(matrix); i++ {
		runSum := 0
		prefix[i+1] = make([]int, len(matrix[i])+1)
		for j := 0; j < len(matrix[i]); j++ {
			runSum += matrix[i][j]
			prefix[i+1][j+1] = runSum + prefix[i][j+1]
		}
	}
	return NumMatrix{prefixMatrix: prefix}
}

func (m *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return m.prefixMatrix[row2+1][col2+1] - m.prefixMatrix[row1][col2+1] - m.prefixMatrix[row2+1][col1] + m.prefixMatrix[row1][col1]
}
