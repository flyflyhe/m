package array

// MaxSubArrayViolence 暴力解法超出时间限制
func MaxSubArrayViolence(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	pre := nums[0]

	for i := 0 ; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			temp := 0
			for x := i; x <= j ; x++ {
				temp += nums[x]
			}
			if temp > pre {
				pre = temp
			}
		}
	}

	return pre
}

func MaxSubArray(nums []int) int  {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] + nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// MaxSubArrayDivide 分治法
func MaxSubArrayDivide(nums []int) int {
	return get(nums, 0, len(nums) - 1).mSum;
}

func pushUp(l, r Status) Status {
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum + r.lSum)
	rSum := max(r.rSum, r.iSum + l.rSum)
	mSum := max(max(l.mSum, r.mSum), l.rSum + r.lSum)
	return Status{lSum, rSum, mSum, iSum}
}

func get(nums []int, l, r int) Status {
	if l == r {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l + r) >> 1
	lSub := get(nums, l, m)
	rSub := get(nums, m + 1, r)
	return pushUp(lSub, rSub)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type Status struct {
	lSum, rSum, mSum, iSum int
}