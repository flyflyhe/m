package tree

import (
	"container/list"
	"flyflyhe.com/m/internal/service/num"
	"fmt"
)

type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}

/**
中根序遍历
 */

func QueueMiddleTree(root *TreeNode)  {
	queue := list.New()
	MiddleTreeErgodic(root, queue)

	for e := queue.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func MiddleTreeErgodic(root *TreeNode, queue *list.List) {
	p := root
	if p != nil {
		if p.Left != nil {
			MiddleTreeErgodic(p.Left, queue)
		}
		queue.PushBack(p.Val)
		if p.Right != nil {
			MiddleTreeErgodic(p.Right, queue)
		}
	}
}

/**
	二叉树的中序遍历
 */

func InorderTraversal(root *TreeNode) []int {
	p := root
	var ret []int
	if p != nil {
		if p.Left != nil {
			ret = append(ret, InorderTraversal(p.Left)...)
		}
		ret = append(ret, p.Val)
		if p.Right != nil {
			ret = append(ret, InorderTraversal(p.Right)...)
		}
	}
	return ret
}

func IsBalanced(root *TreeNode) bool {
	left := 0
	if root.Left != nil {
		left = TreeHeight(root.Left)
		if !IsBalanced(root.Left) {
			return false
		}
	}

	right := 0
	if root.Right != nil {
		right = TreeHeight(root.Right)
		if !IsBalanced(root.Right) {
			return false
		}
	}

	if num.AbsInt(int64(left - right)) > 1 {
		return false
	}

	return true
}

func TreeHeight(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 1
	}

	left := 0
	right := 0
	if root.Left != nil {
		left = TreeHeight(root.Left) + 1
	}

	if root.Right != nil {
		right = TreeHeight(root.Right) + 1
	}

	if left > right {
		return left
	} else {
		return right
	}
}

func IsBalanced2(root *TreeNode) bool {
	left := TreeHeight(root.Left)
	right := TreeHeight(root.Right)
	if num.AbsInt(int64(left - right)) > 1 {
		return false
	}

	if !IsBalanced(root.Left) {
		return false
	}

	if !IsBalanced(root.Right) {
		return false
	}

	return true
}

func TreeHeight2(root *TreeNode) int  {
	if root == nil {
		return 0
	} else {
		left := TreeHeight2(root.Left) + 1
		right := TreeHeight2(root.Right) + 1
		if left > right {
			return left
		} else {
			return right
		}
	}
}
