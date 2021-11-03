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

func RemoveInvalidParentheses(s string) []string {
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

	brackets, _ := checkStr(s)

	//fmt.Println(string(brackets))
	//fmt.Println(s)
	var ret []string
	dfs([]byte(s), brackets, &ret, len(brackets), checkStr)

	var newRet []string

	for _, v := range ret {
		if len(newRet) == 0 {
			newRet = append(newRet, v)
		} else {
			for _, v2 := range newRet {
				if v2 == v {
					goto next
				}
			}
			newRet = append(newRet, v)
		}
		next:
	}

	return newRet
}

func dfs(s []byte, brackets []byte, ret *[]string, deep int, checkStr func(s string) ( []byte, int))  {

	if deep == 0 {
		newS := string(s)
		a, _ := checkStr(newS)
		if len(a) == 0 {
		  	*ret = append(*ret, newS)
		}
		return
	}

	path := make([]byte, len(s))
	for i := 0; i < len(brackets); i++ {
		for j := 0; j < len(s); j++ {
			if s[j] == brackets[i] {
				copy(path, s)
				path := append(path[:j], path[j+1:]...)
				var newBrackets []byte
				if i + 1 < len(brackets) {
					newBrackets = brackets[i+1:]
				}
				//fmt.Println("deep", deep, "i", i, "j", j, "前", string(path), "s", string(s))
				dfs(path, newBrackets, ret, deep-1, checkStr)
			}
		}
	}
}

/**
42: 接雨水
 */

func Trap(height []int) int {
	var stack []int
	length := len(height)
	area := 0
	for i := 0; i < length; i++ {
		for len(stack) > 0 &&  height[stack[len(stack) - 1]] < height[i] {
			top := stack[len(stack)-1] //小于新加入的值
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1] //left >= top
			curWidth := i - left - 1
			curHeight := min(height[left], height[i]) - height[top]
			area += curWidth * curHeight
		}
		stack = append(stack, i)
	}

	return area
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}