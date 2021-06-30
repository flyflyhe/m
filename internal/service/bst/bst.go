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

func (node *Node) Search(v int) *Node {
	if node == nil {
		return nil
	}

	if node.Value > v {
		return node.Left.Search(v)
	} else if node.Value < v {
		return node.Right.Search(v)
	} else {
		return node
	}
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

func (node *Node) Delete(v int) bool {
	needDeleteNode := node.Search(v)
	if needDeleteNode == nil {
		return false
	}

	//左子节点右子节点都为空直接删除
	if needDeleteNode.Left == nil && needDeleteNode.Right == nil {
		if needDeleteNode.Parent.Left.Value == needDeleteNode.Value {
			needDeleteNode.Parent.Left = nil
		} else {
			needDeleteNode.Parent.Right = nil
		}
	} else if needDeleteNode.Right == nil { //只有左节点
		needDeleteNode.Left.Parent = needDeleteNode.Parent
		needDeleteNode.Parent.Left = needDeleteNode.Left
	} else if needDeleteNode.Left == nil { //只有右节点 直接删除此节点 此节点的右子树作为新的
		needDeleteNode.Right.Parent = needDeleteNode.Parent
		needDeleteNode.Parent.Right = needDeleteNode.Right
	} else { //左右节点都不为空 使用前驱节点
		temp := needDeleteNode.Left
		for temp.Right != nil {
			temp = temp.Right
		}

		node.Delete(temp.Value)
		needDeleteNode.Value =  temp.Value
	}

	return true
}

func (node *Node) MiddleRoot() {
	//先输出左叶子节点
	if node.Left != nil {
		node.Left.MiddleRoot()
	}

	//输出本身
	fmt.Print("->", node.Value)

	//输出右叶子
	if node.Right != nil {
		node.Right.MiddleRoot()
	}
}