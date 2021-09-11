package array

import "log"

func TwoSum(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func chalkReplacer(chalk []int, k int) int {
	oneLoop := 0
	for i := 0; i < len(chalk); i++ {
		if k < chalk[i] {
			return i
		}
		oneLoop += chalk[i]
		k -= chalk[i]
	}
	return chalkReplacer(chalk, k % oneLoop)
}

func searchRange(nums []int, target int) []int {
	ret := []int{-1, -1}
	for i, v := range nums {
		if v == target {
			if ret[0] == -1 {
				ret[0] = i
				ret[1] = i
			} else {
				ret[1] = i
			}
		}
	}
	return ret
}

func SearchRange1(nums []int, target int) []int {
	ret := []int{-1, -1}
	l := -1
	r := len(nums)

	if r == 0 {
		return ret
	}

	for l + 1 != r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid
		} else {
			r = mid
		}
	}

	if r >= len(nums) || nums[r] != target {
		return ret
	}

	ret[0] = r
	ret[1] = r
	for {
		r++
		if r >= len(nums) || nums[r] > target {
			break
		}
		if nums[r] == target {
			ret[1] = r
		}
	}

	return ret
}

//
//BlueRedBinarySearch
//找到第一个 >= target的 index 条件 nums[mid] < target 返回r
//找到第一个 < target的 index 条件 nums[mid] < target 返回l
//找到第一个 > target的 index 条件 nums[mid] <= target 返回r
//找到第一个 <= target的 index 条件 nums[mid] <= target 返回l
///**
func BlueRedBinarySearch(nums []int, target int) int  {
	l := -1
	r := len(nums)

	for l + 1 != r {
		mid := (l + r) / 2
		log.Println("l", l, "r", r, "m", mid)
		if nums[mid] < target {
			l = mid
		} else {
			r = mid
		}
	}
	return r
}