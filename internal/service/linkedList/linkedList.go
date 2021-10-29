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

/**
使用递归 从尾部输出链表
 */

func reversePrint(head *ListNode) (ret []int) {

	var reverse func(*ListNode, *[]int)

	reverse = func(listNode *ListNode, ints *[]int) {
		if listNode == nil {
			return
		}

		reverse(listNode.Next, ints)
		*ints = append(*ints, listNode.Val)
	}

	reverse(head, &ret)
	return
}

/**
旋转链表
 */

func RotateRight(head *ListNode, k int) *ListNode {
	p1 := head
	p2 := head
	length := 0
	for p1 != nil {
		length++
		p1 = p1.Next
	}

	if length <= 1 {
		return head
	}

	k = k % length

	if k == 0 {
		return head
	}

	p1 = head

	for i := 0; i < length; i++ {
		if i + k == length {
			break
		}
		p1 = p1.Next
	}

	for i := 0; i < length - k - 1; i++ {
		p2 = p2.Next
	}
	p2.Next = nil

	newHead := p1

	for p1.Next != nil {
		p1 = p1.Next
	}
	p1.Next = head

	return newHead
}

func PrintList(node *ListNode) (ret []int)  {
	for node != nil {
		ret = append(ret, node.Val)
		node = node.Next
	}
	return
}