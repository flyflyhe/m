package array

import "fmt"

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
		for j := 0; j < i; j++ {
			if nums[j] > nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}

	return nums
}

func SortNumQuack(nums []int) {
	if len(nums) < 2 {
		return
	}
	m := nums[0] //用首位元素分割
	r := len(nums) - 1
	left := false
	for l := 0; l < r; l++ {
		if !left {
			if nums[r] < m{
				nums[l], nums[r] = nums[r], nums[l]
				left = true
			} else {
				l--
				r--
			}
		} else {
			if nums[l] > m{
				nums[l], nums[r] = nums[r], nums[l]
				left = false
				l--
			}
		}
	}
	nums[r] = m
	fmt.Println(nums)
	SortNumQuack(nums[0:r])
	SortNumQuack(nums[r+1:])
}

func SortNum(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	l := 0
	r := len(nums) - 1
	m := nums[0]

	for l < r {
		for ;l < r; r-- {
			if nums[r] < m {
				nums[l] = nums[r]
				break
			}
		}
		for ; l < r; l++ {
			if nums[l] > m {
				nums[r] = nums[l]
				break
			}
		}
	}

	left := nums[:l]
	right := nums[r+1:]
	fmt.Println(left)
	fmt.Println(right)
	ret := append(SortNum(left), m)

	return append(ret, SortNum(right)...)
}


