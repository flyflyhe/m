package array

import (
	"fmt"
)

func MinAbsoluteSumDiffViolence(nums1 []int, nums2 []int) int {
	var result int64

	minAbsoluteNum := -1
	index1 := 0
	index2 := 0

	for i := 0; i < len(nums1); i++ {
		absoluteNum := abs(nums1[i] - nums2[i])
		result += int64(absoluteNum)

		for j := 0 ; j < len(nums2); j++ {
			if i == j {
				continue
			}

			//换位后一定要比换位之前小  否则没有意义
			if abs(nums1[i] - nums2[j]) > abs(nums1[j] - nums2[j]) {
				continue
			}

			if minAbsoluteNum == -1 {
				minAbsoluteNum = abs(nums1[j] - nums2[j]) - abs(nums1[i] - nums2[j])
				index1 = i
				index2 = j
			}  else {
				tmp := abs(nums1[j] - nums2[j]) - abs(nums1[i] - nums2[j])
				if tmp > minAbsoluteNum {
					minAbsoluteNum = tmp
					index1 = i
					index2 = j
				}
			}
		}
	}

	fmt.Println("i", index1, "j", index2)

	if minAbsoluteNum != -1 {
		result -= int64(minAbsoluteNum)
	}

	return int(result % int64(1000000000 + 7))
}

func abs(n int) int {
	y := n >> 31
	return (n ^ y) - y
}

func abs64(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}
