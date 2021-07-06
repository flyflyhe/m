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

// RemoveElementGood 可以保证数组顺序
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

// RemoveElementRight 此方法效率高 但不能保证数组的顺序
func RemoveElementRight(nums []int, val int) int  {
	right := len(nums) - 1
	//左指针要小于有指针
	for left := 0; left <= right; left++ {
		if nums[left] == val { //如果左侧位置值相等
			nums[left] = nums[right]  //把右侧值 放到左侧
			left--  //为什么要--  因为交换完的数据也要 验证是否相等
			right--
		}
	}

	return right + 1 //为什么+1  因为循环里面最后面减1
}