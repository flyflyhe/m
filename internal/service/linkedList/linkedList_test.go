package linkedList

import (
	"fmt"
	"testing"
)

func TestRotateRight(t *testing.T) {
	head := &ListNode{Next: &ListNode{Next: &ListNode{Next: &ListNode{Next: &ListNode{Next: nil, Val: 5}, Val: 4}, Val: 3}, Val: 2}, Val: 1}
	fmt.Println(PrintList(RotateRight(head, 2)))
}
