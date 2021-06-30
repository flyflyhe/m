package bst

import "fmt"

type Node struct {
	Value int
	Parent *Node
	Left *Node
	Right *Node
}

func GenerateBst(list []int) *Node {
	node := InitNode()
	for _, v := range list {
		node.Insert(v)
	}

	return node
}

func InitNode() *Node  {
	node := &Node{Value: -1}
	return node
}

func (node *Node) Insert(v int) bool {
	if node.Value == -1 {
		node.Value = v
	} else if node.Value > v {
		if node.Left != nil {
			node.Left.Insert(v)
		} else {
			node.Left = InitNode()
			node.Left.Value = v
			node.Left.Parent = node
		}
	} else if node.Value < v {
		if node.Right != nil {
			node.Right.Insert(v)
		} else {
			node.Right = InitNode()
			node.Right.Value = v
			node.Right.Parent = node
		}
	} else {
		return false
	}
	return true
}

func (node *Node) MiddleRoot() {
	//先输出左叶子节点
	if node.Left != nil {
		node.Left.MiddleRoot()
	}

	//输出本身
	fmt.Print(node.Value, ",")

	//输出右叶子
	if node.Right != nil {
		node.Right.MiddleRoot()
	}
}