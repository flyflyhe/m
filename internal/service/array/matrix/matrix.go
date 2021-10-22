package matrix

func Rotate(matrix [][]int)  {
	rows := len(matrix)
	cols := len(matrix[0])
	matrixNew := make([][]int, rows)
	for i, _ := range matrixNew {
		matrixNew[i] = make([]int, cols)
	}
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			matrixNew[j][cols-i-1] = matrix[i][j]
		}
	}

	copy(matrix, matrixNew)
}

func Rotate2(matrix [][]int)  {
	rows := len(matrix)
	cols := len(matrix[0])

	//水平翻转
	for i := 0; i < rows / 2; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j], matrix[rows-i-1][j] = matrix[rows-i-1][j], matrix[i][j]
		}
	}

	//对角线翻转
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == j {
				break
			}
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
