package array

import "testing"

func TestMinPairSum(t *testing.T) {
	nums := []int{3,5,2,3}

	r := MinPairSum(nums)

	if r != 7 {
		t.Errorf("ææ%dè·å¾%d", 7, r)
	}
}
