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

func LargestRectangleArea(heights []int) int {
	stack := make([]int, 0)
	area := 0
	for i, v := range heights {
		for len(stack) > 0 && v < heights[stack[len(stack) - 1]] {
			hIndex := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			startIndex := -1
			if len(stack) > 0 {
				startIndex = stack[len(stack)-1]
			}
			h := heights[hIndex]
			w := i - startIndex - 1
			tmpArea := h * w
			if tmpArea > area {
				area = tmpArea
			}
		}
		stack = append(stack, i)
	}

	for len(stack) > 0 {
		hIndex := stack[len(stack) - 1]
		stack  = stack[:len(stack) - 1]
		startIndex := -1
		if len(stack) > 0 {
			startIndex = stack[len(stack)-1]
		}
		h := heights[hIndex]
		w := len(heights) - startIndex - 1
		tmpArea := h * w
		if tmpArea > area {
			area = tmpArea
		}
	}
	return area
}

/**
删除无用括号
 */

func removeInvalidParentheses(s string) []string {
	checkStr := func(s string) (stack []byte, bracketCounter int) {
		for i := 0; i < len(s); i++ {
			if s[i] != '(' && s[i] != ')' {
				continue
			}
			bracketCounter++
			if len(stack) > 0 && s[i] != stack[len(stack) - 1] && s[i] == ')' { //凑对 弹出
				stack = stack[:len(stack) - 1]
			} else {
				stack = append(stack, s[i])
			}
		}
		return
	}
	brackets, num := checkStr(s)
	var ret []string
	if num == len(brackets) {
		ret = append(ret, "")
		return ret
	}
	//大于0 说明需要删除的对
	if len(brackets) > 0 {
		var path []byte
		path = append(path, s[:]...)
		dfs([]byte(s), brackets, path, &ret, checkStr)
	}
	return ret
}

func dfs(s []byte,  brackets []byte, path []byte, ret *[]string, checkStr func(s string) ([]byte, int))  {
	brackets, _ = checkStr(string(path))
	if len(brackets) == 0 {
		*ret = append(*ret, string(path))
		return
	}

	for i := 0; i < len(brackets); i++ {
		for j := 0; j < len(s); j++ {
			if s[j] == brackets[i] {
				path = append(path[:j], path[j+1:]...)
				dfs(s, brackets, path, ret, checkStr)
			}
		}
	}
}