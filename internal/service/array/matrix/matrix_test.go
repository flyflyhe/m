package matrix

import (
	"fmt"
	"testing"
)

var matrix [][]int

func init()  {
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
