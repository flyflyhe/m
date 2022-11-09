package matrix

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var matrix [][]int

func init() {
	matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},

		//水平翻转
		//{7, 8, 9},
		//{4, 5, 6},
		//{1, 2, 3},

		//{7, 3, 1},
		//{8, 5, 2},
		//{9, 6, 3},
	}
}

func TestRotate(t *testing.T) {
	Rotate(matrix)
	fmt.Println(matrix)
}

func TestRotate2(t *testing.T) {
	Rotate2(matrix)
	fmt.Println(matrix)
}

func TestSearchMatrix(t *testing.T) {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}}
	fmt.Println(SearchMatrix(matrix, 5))
	matrix = [][]int{{-1}, {-1}}
	fmt.Println(SearchMatrix(matrix, 21))
}

func TestMaximalRectangle(t *testing.T) {
	fmt.Println(MaximalRectangle([]string{"01", "01"}))
}

func TestMaxCount(t *testing.T) {
	fmt.Println(MaxCount(3, 3, [][]int{{2, 2}, {3, 3}}))
}

func TestIsValidSudoku(t *testing.T) {
	matrix := [][]byte{
		[]byte{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
		[]byte{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
		[]byte{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
		[]byte{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
		[]byte{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
		[]byte{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
	}
	fmt.Println(IsValidSudoku(matrix))

}

func Test_orderOfLargestPlusSign(t *testing.T) {
	ans := orderOfLargestPlusSign(2, [][]int{{0, 0}, {0, 1}, {1, 0}})
	assert.Equal(t, 1, ans)
}
