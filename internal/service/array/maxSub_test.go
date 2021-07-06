package array

import "testing"

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}

	r := MaxSubArray(nums)

	if r != 6 {
		t.Errorf("期望%d获取%d", 6, r)
	}

	nums = []int{5,4,-1,7,8}
	r = MaxSubArray(nums)

	if r != 23 {
		t.Errorf("期望%d获取%d", 23, r)
	}

	nums = []int{-2, 1}
	r = MaxSubArray(nums)

	if r != 1 {
		t.Errorf("期望%d获取%d", 1, r)
	}
}
