package matrix

import (
	"flyflyhe.com/m/internal/service/stack"
	"fmt"
)

func Rotate(matrix [][]int) {
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

func Rotate2(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])

	//水平翻转
	for i := 0; i < rows/2; i++ {
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
		if direction == 0 { //向右
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

		} else { //向下
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

func MaximalRectangle(matrix []string) int {
	length := len(matrix)
	if length == 0 {
		return 0
	}
	area := 0
	height := make([]int, len(matrix[0]))
	for i := 0; i < length; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			height[j] = int(matrix[i][j] - '0')
			if height[j] == 1 {
				for x := i - 1; x >= 0; x-- {
					if matrix[x][j] == '1' {
						height[j]++
					} else {
						break
					}
				}
			}
		}
		fmt.Println(height)
		tmpArea := stack.LargestRectangleArea(height)
		if tmpArea > area {
			area = tmpArea
		}
	}
	return area
}

func isSelfCrossing(distance []int) bool {
	m := make(map[[2]int]int)

	m[[2]int{0, 0}] = 0
	direction := [4][]int{
		{0, 1},
		{-1, 0},
		{0, -1},
		{1, 0},
	}
	row := 0
	col := 0
	for i := 0; i < len(distance); i++ {
		n := distance[i]
		for j := 0; j < n; j++ {
			di := i % 4
			row += direction[di][0]
			col += direction[di][1]

			if _, ok := m[[2]int{row, col}]; ok {
				return true
			} else {
				m[[2]int{row, col}] = 0
			}
		}
	}

	return false
}

/**
:598
范围求和
*/

func MaxCount(m int, n int, ops [][]int) int {
	a := m
	b := n
	for i := 0; i < len(ops); i++ {
		if ops[i][0] < a {
			a = ops[i][0]
		}

		if ops[i][1] < b {
			b = ops[i][1]
		}
	}

	return a * b
}

func IsValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		m := make(map[byte]int)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if _, ok := m[board[i][j]]; ok {
					return false
				} else {
					m[board[i][j]] = 1
				}
			}
		}
	}

	for i := 0; i < 9; i++ {
		m := make(map[byte]int)
		for j := 0; j < 9; j++ {
			if board[j][i] != '.' {
				if _, ok := m[board[j][i]]; ok {
					return false
				} else {
					m[board[j][i]] = 1
				}
			}
		}
	}

	matrix := [][2][2]int{
		[2][2]int{[2]int{0, 0}, [2]int{2, 2}},
		[2][2]int{[2]int{0, 3}, [2]int{2, 5}},
		[2][2]int{[2]int{0, 6}, [2]int{2, 8}},
		[2][2]int{[2]int{3, 0}, [2]int{5, 2}},
		[2][2]int{[2]int{3, 3}, [2]int{5, 5}},
		[2][2]int{[2]int{3, 6}, [2]int{5, 8}},
		[2][2]int{[2]int{6, 0}, [2]int{8, 2}},
		[2][2]int{[2]int{6, 3}, [2]int{8, 5}},
		[2][2]int{[2]int{6, 6}, [2]int{8, 8}},
	}

	for _, v := range matrix {
		var m [10]int8
		for i := v[0][0]; i <= v[1][0]; i++ {
			for j := v[0][1]; j <= v[1][1]; j++ {
				val := board[i][j]
				if val == '.' {
					continue
				}
				if m[val-'0'] == 0 {
					m[val-'0'] = 1
				} else {
					return false
				}
			}
		}
	}

	return true
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	rows := make([]int, len(grid))
	if len(rows) == 0 {
		return -1
	}
	cols := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			rows[i] = max(rows[i], grid[i][j])
			cols[j] = max(cols[j], grid[i][j])
		}
	}

	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if rows[i] < cols[j] {
				ans += rows[i] - grid[i][j]
			} else {
				ans += cols[j] - grid[i][j]
			}
		}
	}

	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

//:764 最大加号标志

func orderOfLargestPlusSign(n int, mines [][]int) int {
	minesMap := make(map[[2]int]struct{})
	grid := make([][]int, n)

	for _, v := range mines {
		minesMap[[2]int{v[0], v[1]}] = struct{}{}
	}

	for k, _ := range grid {
		grid[k] = make([]int, n)
		for i := 0; i < n; i++ {
			if _, ok := minesMap[[2]int{k, i}]; ok {
				grid[k][i] = 0
			} else {
				grid[k][i] = 1
			}
		}
	}

	ans := -1
	x := [4][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}
	fmt.Println(grid)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			min := n
			for _, v := range x {
				t := 0
				r := i
				c := j
				for grid[r][c] == 1 {
					t++
					r = r + v[0]
					c = c + v[1]
					fmt.Println("r", r, "c", c)
					if r < 0 || c < 0 || r == n || c == n {
						break
					}
				}
				if t < min {
					min = t
				}
			}
			if min > ans && min != n {
				ans = min
			}
		}
	}
	return ans
}
