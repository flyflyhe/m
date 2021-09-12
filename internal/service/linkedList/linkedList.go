package linkedList

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		point := head //point 指向当前节点
		head = head.Next //head 立马指向下一节点 不能放后面 因为point.next 会变化

		point.Next = prev
		prev = point
	}

	return prev
}