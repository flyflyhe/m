package array

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
	l := 0
	r := len(nums) - 1

	for ; l <= r; {
		mid := l + r + 1 >> 1
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	if r >= 0 {
		if nums[r] == target {
			l = r
		} else {
			if r == 0 || nums[r - 1] != target {
				return ret
			}
			l = r - 1
			r = l
		}
	} else {
		return ret
	}

	for {
		if l < 0 ||  nums[l] < target {
			break
		}
		if nums[l] == target {
			ret[0] = l
		}
		l--
	}

	for {
		if r >= len(nums) || nums[r] > target {
			break
		}
		if nums[r] == target {
			ret[1] = r
		}
		r++
	}

	return ret
}