package array

import "testing"

func TestMinAbsoluteSumDiffViolence(t *testing.T) {
	nums1 := []int{1,7,5}
	nums2 := []int{2,3,5}

	result := MinAbsoluteSumDiffViolence(nums1, nums2)

	if result != 3 {
		t.Errorf("期望%d获得%d", 3, result)
	}

	nums1 = []int{1,10,4,4,2,7}
	nums2 = []int{9,3,5,1,7,4}

	result = MinAbsoluteSumDiffViolence(nums1, nums2)
	if result != 20 {
		t.Errorf("期望%d获得%d", 20, result)
	}
}
