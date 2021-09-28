package recursion

import "fmt"

func IsPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}

	if n == 1 {
		return true
	}

	if n % 3 != 0 {
		return false
	}

	return IsPowerOfThree(n / 3)
}

func reverseString(s []byte)  {
	l := 0
	r := len(s) - 1
	for l < r {
		temp := s[l]
		s[l] = s[r]
		s[r] = temp

		r--
		l++
	}

	return
}

type ListNode struct {
	Val int
	Next *ListNode
}

func ReversePrint(head *ListNode) []int {
	var arr []int
	for head != nil {
		fmt.Println(head.Val)
		arr = append(arr, head.Val)
		head = head.Next
	}

	l := 0
	r := len(arr) - 1
	for l < r {
		temp := arr[l]
		arr[l] = arr[r]
		arr[r] = temp

		r--
		l++
	}

	return arr
}