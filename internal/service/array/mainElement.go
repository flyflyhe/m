package array

func majorityElementViolence(nums []int) int {
	m := make(map[int]int)
	majorElement := -1

	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}

	for e, v := range m{
		if v > len(nums) / 2 {
			majorElement = e
		}
	}

	return majorElement
}

func majorityElement(nums []int) int {
	element := nums[0]
	elementLen := 1
	nums = SortNumsPop(nums)
	for i := 1; i < len(nums); i++ {
		if element == nums[i] {
			elementLen++
		} else {
			if elementLen > len(nums) / 2 {
				return element
			}
			element = nums[i]
			elementLen = 1
		}
	}

	if elementLen > len(nums) / 2 {
		return element
	}

	return -1
}

func SortNumsPop(nums []int) []int  {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}

	return nums
}
