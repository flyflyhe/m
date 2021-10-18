package heap

import (
	"math"
)

type Node struct {
	Key int
	Val int
	Left *Node
	Right *Node
	Parent *Node
}

type MaxHeap struct {
	root *Node
	height int
	num int
}

func (this *MaxHeap) GetRoot() *Node {
	return this.root
}

func (this *MaxHeap) Insert(Key int, val int)  {
	if this.root == nil {
		this.root = &Node{Key: Key}
		this.height = 1
		this.num = 1
	} else {
		if this.num == int(math.Pow(float64(2), float64(this.height-1))) {
			this.height++
			insert(this.root, Key, val, this.height - 1)
		} else {
			insert(this.root, Key, val, this.height - 1)
		}
		this.num++
	}
}

func (this *MaxHeap) Remove() (int, int)  {
	if this.num == 0 {
		return -1, -1
	}
	root := this.root
	key := root.Key
	val := root.Val

	s := PreOrder(root)
	tmpKey := s[len(s)-1]
	head := &Node{Left: root}
	delKey(head, tmpKey)
	root.Key = tmpKey
	rebuildTop(root)

	this.num--
	if this.num == 0 {
		this.root = nil
	}
	return key, val
}

func insert(root *Node, Key int, val int, height int) bool  {
	if height == 1 {
		if root.Left == nil {
			root.Left = &Node{Key: Key,Parent: root, Val: val}
			rebuild(root.Left)
		} else if root.Right == nil {
			root.Right = &Node{Key: Key,Parent: root, Val: val}
			rebuild(root.Right)
		} else {
			return false
		}

		return true
	}
	if insert(root.Left, Key, val, height-1) {
		return true
	}
	return insert(root.Right, Key, val, height-1)
}

func rebuildTop(root *Node)  {
	if root == nil {
		return
	}
	left := root.Left
	right := root.Right
	if left != nil && right != nil {
		if left.Key > right.Key {
			if left.Key > root.Key {
				left.Key, root.Key = root.Key, left.Key
				left.Val, root.Val = root.Val, left.Val
				rebuildTop(left)
			}
		} else {
			if right.Key > root.Key {
				right.Key, root.Key = root.Key, right.Key
				right.Val, root.Val = root.Val, right.Val
				rebuildTop(right)
			}
		}
	} else if left != nil {
		if left.Key > root.Key {
			left.Key, root.Key = root.Key, left.Key
			left.Val, root.Val = root.Val, left.Val
			rebuildTop(left)
		}
	} else if right != nil {
		if right.Key > root.Key {
			right.Key, root.Key = root.Key, right.Key
			right.Val, root.Val = root.Val, right.Val
			rebuildTop(right)
		}
	}
}

func rebuild(root *Node)  {
	if root.Parent == nil {
		return
	}
	if root.Parent.Key < root.Key {
		root.Parent.Key, root.Key = root.Key, root.Parent.Key
		root.Parent.Val, root.Val = root.Val, root.Parent.Val
		rebuild(root.Parent)
	}
}


func delKey(root *Node, key int) {
	p := root
	if p != nil {
		if p.Left != nil {
			if p.Left.Key == key {
				p.Left = nil
			} else {
				delKey(p.Left, key)
			}
		}

		if p.Right != nil {
			if p.Right.Key == key {
				p.Right = nil
			} else {
				delKey(p.Right, key)
			}
		}
	}
}

func PreOrder(root *Node) []int {
	p := root
	var ret []int
	if p != nil {
		ret = append(ret, p.Key)
		ret = append(ret, p.Val)
		if p.Left != nil {
			ret = append(ret, PreOrder(p.Left)...)
		}

		if p.Right != nil {
			ret = append(ret, PreOrder(p.Right)...)
		}
	}
	return ret
}





