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
