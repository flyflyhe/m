package completeBinaryTreeInserter

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CBTInserter struct {
	TreeNodeList  []int
	LastLevelNode int //root 第0层
	Level         int
}

func Constructor(root *TreeNode) CBTInserter {
	if root == nil {
		return CBTInserter{
			TreeNodeList:  []int{},
			LastLevelNode: 0,
			Level:         0,
		}
	}
	l := []int{}

	queue := []*TreeNode{}

	queue = append(queue, root)

	level := -1
	lastLevelNode := 0
	for len(queue) > 0 {
		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			node := queue[i]
			l = append(l, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		level++
		lastLevelNode = qLen
		queue = queue[qLen:]
	}

	return CBTInserter{TreeNodeList: l, Level: level, LastLevelNode: lastLevelNode}
}

func (this *CBTInserter) Insert(val int) int {
	this.TreeNodeList = append(this.TreeNodeList, val)
	if this.Level == 0 && this.LastLevelNode == 0 { //树为空 此种情况应该不会出现 因为无法返回数据
		this.Level = 0
		this.LastLevelNode = 1
		return -1
	} else if this.Level == 0 && this.LastLevelNode == 1 {
		this.Level++
		this.LastLevelNode = 1
		return this.TreeNodeList[0]
	} else if int(math.Pow(float64(2), float64(this.Level))) == this.LastLevelNode {
		this.Level++
		this.LastLevelNode = 1
	} else {
		this.LastLevelNode++
	}

	var parentIndex int
	if this.LastLevelNode%2 == 0 { //说明最后插入的是右节点
		parentIndex = (len(this.TreeNodeList) - 2) / 2
	} else { //左节点
		parentIndex = (len(this.TreeNodeList) - 1) / 2
	}

	return this.TreeNodeList[parentIndex]
}

func (this *CBTInserter) Get_root() *TreeNode {
	return nil
}
