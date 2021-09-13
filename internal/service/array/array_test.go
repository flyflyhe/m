package array

import "testing"

func TestBinarySearch(t *testing.T) {
	nums := []int{-1,0,3,3,3,3,3,3,3,5,9,12}

	t.Log(BinarySearch(nums, 3))
}
