package stack

import (
	"fmt"
	"testing"
)

func TestValidateStackSequences(t *testing.T) {
	pushed := []int{1,2,3,4,5}
	popped := []int{4,5,3,2,1}
	fmt.Println(ValidateStackSequences(pushed, popped))
	fmt.Println(ValidateStackSequences(pushed, []int{4,3,5,1,2}))
	fmt.Println(ValidateStackSequences([]int{1, 0}, []int{1, 0}))
}

func TestLargestRectangleArea(t *testing.T) {
	fmt.Println(LargestRectangleArea([]int{2,4}))
}

/**
40 (
41 )
 */
func TestRemoveInvalidParentheses(t *testing.T) {
	fmt.Println(RemoveInvalidParentheses("()())()"))
	fmt.Println(RemoveInvalidParentheses("(a)())()"))
	fmt.Println(RemoveInvalidParentheses(")("))
}

func TestTrap(t *testing.T)  {
	fmt.Println(Trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
}