package array

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{-1,0,3,3,3,3,3,3,3,5,9,12}

	t.Log(BinarySearch(nums, 3))
}

func TestGetRow(t *testing.T) {

	fmt.Println(GetRow(3))
	fmt.Println(GetRow(0))
	fmt.Println(GetRow(1))
}
