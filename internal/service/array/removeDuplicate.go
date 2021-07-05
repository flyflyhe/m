package array

import "fmt"

func RemoveDuplicate(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last, finder := 0, 0//快慢指针
	for last < len(nums)-1 {
		for nums[finder] == nums[last] {
			finder++
			if finder == len(nums) {
				return last + 1
			}
		}
		nums[last+1] = nums[finder]//last+1位置已经比对过 可以复用 要满足O(1)空间复杂度
		last++
	}
	return last + 1
}

func RemoveElement(nums []int, v int) int  {
	slow :=  0

	fmt.Println("============================")
	for fast := 0 ; fast < len(nums) ; fast++ {
		for nums[fast] == v {
			if fast == len(nums) - 1 {
				break
			}
			fast++
		}

		fmt.Println(nums)
		fmt.Println("slow", slow, "fast", fast)

		if nums[fast] != v {
			if slow < fast {
				nums[slow] = nums[fast]
			}
		} else {
			break
		}
		slow++
		fast++
	}

	fmt.Println(nums)

	return slow
}

func RemoveElementGood(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			fmt.Println(nums)
			fmt.Println("fast", fast, "slow", slow)
			if slow != fast {
				nums[slow] = nums[fast]
			}
			slow++
		}
	}
	return slow

}