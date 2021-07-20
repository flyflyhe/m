package array

import (
	"fmt"
	"sort"
)

func MinPairSum(nums []int) int {
	sort.Ints(nums)
	fmt.Println(nums)
	length := len(nums)

	maxPair := 0
	for i := 0; i < length / 2; i++ {
		temp := nums[i] + nums[length - i - 1]
		if i == 0 {
			maxPair = temp
		} else {
			if temp > maxPair {
				maxPair = temp
			}
		}
	}

	return maxPair
}