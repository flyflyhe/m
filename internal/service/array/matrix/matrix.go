package matrix

import "fmt"

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

func SearchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	var binarySearch func(int, int, int, int) (int, int)
	binarySearch = func(row, col, target, direction int) (int, int) {
		ret := []int{0, 0}
		if direction == 0 {//向右
			ret[0] = row
			l := col
			r := cols - 1
			for l < r {
				mid := (l + r + 1) >> 1
				if matrix[row][mid] == target {
					ret[1] = mid
					break
				} else if matrix[row][mid] < target {
					l = mid
				} else {
					r = mid - 1
				}
				ret[1] = r
			}

		} else {//向下
			ret[1] = col
			l := row
			r := rows - 1
			fmt.Println("下", "l", l, "r", r)
			for l < r {
				mid := (l + r + 1) >> 1
				if matrix[mid][col] == target {
					ret[0] = mid
					break
				} else if matrix[mid][col] < target {
					l = mid
				} else {
					r = mid - 1
				}
				ret[0] = r
			}
		}

		return ret[0], ret[1]
	}

	col := 0
	for row := 0; row < rows; row++ {
		x, y := binarySearch(row, col, target, 0)
		if matrix[x][y] == target {
			return true
		}
		x, y = binarySearch(x, y, target, 1)
		if matrix[x][y] == target {
			return true
		}
	}

	return false
}
