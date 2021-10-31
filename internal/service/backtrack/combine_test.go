package backtrack

import (
	"fmt"
	"testing"
)

func TestCombine(t *testing.T) {
	fmt.Println(Combine(4, 2))
	fmt.Println(Combine(1, 1))
}

func TestMaxTwoEvents(t *testing.T) {
	matrix := [][]int{
		//{1,5,3},{1,5,1},{6,6,5},
		//{1,3,2},{4,5,2},{2,4,3},
		//{10,83,53},{63,87,45},{97,100,32},{51,61,16},
		{66,97,90},{98,98,68},{38,49,63},{91,100,42},{92,100,22},{1,77,50},{64,72,97},
	}

	fmt.Println(MaxTwoEvents(matrix))
}
