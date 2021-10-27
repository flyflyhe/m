package array

import (
	"fmt"
	"testing"
)

func TestSortArray(t *testing.T) {
	fmt.Println(SortArray([]int{5,4,3,2,1}))
}

func TestMergeSort(t *testing.T) {
	nums := []int{5,4,3,2,1}
	fmt.Println(MergeSort(nums))
	fmt.Println(nums)
}
