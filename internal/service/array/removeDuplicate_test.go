package array

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicate(t *testing.T) {
	nums := []int{0,0,1,1,1,2,2,3,3,4}

	r := RemoveDuplicate(nums)
	if r != 5 {
		t.Errorf("期望%d获得%d", 5, r)
	}
}

func TestRemoveElement(t *testing.T) {
	nums := []int{3,2,2,3}
	r := RemoveElement(nums, 3)

	fmt.Println(nums)
	if r != 2 {
		t.Errorf("期望%d获得%d", 2, r)
	}

	nums = []int{0,1,2,2,3,0,4,2}
	r = RemoveElement(nums, 2)

	fmt.Println(nums)
	if r != 5 {
		t.Errorf("期望%d获得%d", 5, r)
	}

	r = RemoveElement(nums, 4)
	fmt.Println(nums)
	if r != 7 {
		t.Errorf("期望%d获得%d", 7, r)
	}
}
