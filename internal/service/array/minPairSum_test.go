package array

import "testing"

func TestMinPairSum(t *testing.T) {
	nums := []int{3,5,2,3}

	r := MinPairSum(nums)

	if r != 7 {
		t.Errorf("期望%d获得%d", 7, r)
	}
}
