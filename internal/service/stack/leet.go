package stack

import (
	"container/list"
)

func ValidateStackSequences(pushed []int, popped []int) bool {
	stack := list.New()
	i := 0
	for _, v := range pushed {
		stack.PushBack(v)
		for ele := stack.Back(); ele != nil && ele.Value.(int) == popped[i]; i++{
			stack.Remove(ele)
			ele = stack.Back()
		}
	}

	return stack.Len() == 0
}

/**
使用单调栈
[73,74,75,71,69,72,76,73]
 */

func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	if length == 0 {
		return nil
	}
	ret := make([]int, length)
	stack := make([]int, 1)
	stack[0] = 0
	for i := 1; i < length; i++ {
		stackIndex := stack[len(stack) - 1]
		if temperatures[i] <= temperatures[stackIndex] {
			stack = append(stack, i)
		} else {
			for temperatures[i] > temperatures[stackIndex] { //新入栈大于栈顶
				ret[stackIndex] = i - stackIndex
				stack = stack[:len(stack) - 1]
				if len(stack) == 0 {
					break
				}
				stackIndex = stack[len(stack) - 1]
			}
			stack = append(stack, i)
		}
	}
	return ret
}