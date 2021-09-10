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