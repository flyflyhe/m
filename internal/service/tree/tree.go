package tree

import (
	"container/list"
	"flyflyhe.com/m/internal/service/num"
	"fmt"
	"log"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pruneTree(root *TreeNode) *TreeNode {
	tapNode(root)
	DeleteNode(root)
	if root.Val == -1 {
		return nil
	}
	return root
}

func DeleteNode(node *TreeNode) {
	if node == nil {
		return
	}

	if node.Left != nil && node.Left.Val == -1 {
		node.Left = nil
	} else {
		DeleteNode(node.Left)
	}

	if node.Right != nil && node.Right.Val == -1 {
		node.Right = nil
	} else {
		DeleteNode(node.Right)
	}
}

func tapNode(node *TreeNode) bool {
	if node == nil {
		return true
	}
	if node.Val == 1 {
		tapNode(node.Left)
		tapNode(node.Right)
		return false
	} else {
		l := tapNode(node.Left)
		r := tapNode(node.Right)

		if l && r {
			node.Val = -1
			return true
		} else {
			return false
		}
	}
}

/**
中根序遍历
*/

func QueueMiddleTree(root *TreeNode) {
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

func PreOrder(root *TreeNode) []int {
	p := root
	var ret []int
	if p != nil {
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

	if num.AbsInt(int64(left-right)) > 1 {
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
	if num.AbsInt(int64(left-right)) > 1 {
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

func TreeHeight2(root *TreeNode) int {
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

func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{}
	log.Println(preorder, inorder)
	v := preorder[0] //根
	root.Val = v
	index := SearchIndex(inorder, v)
	log.Println(index, index)
	left := inorder[:index]    //左子树
	right := inorder[index+1:] //右子树

	if len(left) > 0 {
		root.Left = BuildTree(preorder[1:len(left)+1], left)
	}

	if len(right) > 0 {
		root.Right = BuildTree(preorder[len(left)+1:], right)
	}

	return root
}

func SearchIndex(arr []int, val int) int {
	for i := 0; i < len(arr); i++ {
		if val == arr[i] {
			return i
		}
	}
	return -1
}

func KthSmallest(root *TreeNode, k int) int {
	var s []int
	var preFunc func(root *TreeNode)
	preFunc = func(root *TreeNode) {
		if root != nil {
			s = append(s, root.Val)
			preFunc(root.Left)
			preFunc(root.Right)
		}
	}

	preFunc(root)
	return GetK(s, k)
}

func GetK(s []int, k int) int {
	InsertSort(s)
	for i, v := range s {
		if i == k-1 {
			return v
		}
	}

	return 0
}

func InsertSort(st []int) {
	if len(st) <= 1 {
		return
	}
	for i := 1; i < len(st); i++ {
		back := st[i]
		j := i - 1
		for j >= 0 && back < st[j] {
			st[j+1] = st[j]
			j--
		}
		st[j+1] = back
	}
	return
}

func VerifyPostorder(postorder []int) bool {
	return true
}

func PostorderTraversalRecursion(root *TreeNode) []int {
	var ret []int

	var recursion func(*TreeNode, *[]int)
	recursion = func(node *TreeNode, ret *[]int) {
		if node != nil {
			recursion(node.Left, ret)
			recursion(node.Right, ret)
			*ret = append(*ret, node.Val)
		}
	}

	recursion(root, &ret)

	return ret
}

func PostorderTraversal(root *TreeNode) (res []int) {
	var stk []*TreeNode
	var prev *TreeNode //prev 防止有右子树的情况下出现死循环
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stk = append(stk, root)
			root = root.Right
		}
	}
	return
}

func PreorderTraversal(root *TreeNode) (res []int) {
	var stk []*TreeNode
	for root != nil || len(stk) > 0 {
		for root != nil {
			res = append(res, root.Val)
			stk = append(stk, root)
			root = root.Left
		}

		root = stk[len(stk)-1].Right //不需要在回到根节点
		stk = stk[:len(stk)-1]
	}

	return
}

/**
树的层序遍历
*/

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	arr := make([][]int, TreeHeight(root))
	queue := list.New()
	level := 0
	queue.PushBack(root)
	for queue.Len() > 0 {
		eleSlice := []*list.Element{}
		for queue.Len() > 0 {
			ele := queue.Front()
			eleSlice = append(eleSlice, ele)
			queue.Remove(ele)
		}

		for i := 0; i < len(eleSlice); i++ {
			node := eleSlice[i].Value.(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}

			if arr[level] == nil {
				arr[level] = make([]int, 0)
			}
			arr[level] = append(arr[level], node.Val)
		}
	}

	return arr
}

/**
对称二叉树 先根序遍历
*/

func IsSymmetric(root *TreeNode) bool {
	left := LeftOrder(root.Left)
	right := RightOrder(root.Right)
	fmt.Println("left", left)
	fmt.Println("right", right)
	for i, v := range left {
		if right[i] != v {
			return false
		}
	}

	return true
}

func RightOrder(root *TreeNode) []int {
	p := root
	if p == nil {
		return []int{-1}
	}
	var ret []int
	ret = append(ret, p.Val)
	ret = append(ret, RightOrder(p.Right)...)
	ret = append(ret, RightOrder(p.Left)...)
	return ret
}

func LeftOrder(root *TreeNode) []int {
	p := root
	if p == nil {
		return []int{-1}
	}
	var ret []int
	ret = append(ret, p.Val)
	ret = append(ret, LeftOrder(p.Left)...)
	ret = append(ret, LeftOrder(p.Right)...)
	return ret
}

type Node struct {
	Val      int
	Children []*Node
}

func bfs(root *Node) int {
	if root == nil {
		return 0
	}
	s := make([]*Node, 0)
	s = append(s, root)
	ans := 0
	for len(s) > 0 {
		temp := s
		s = []*Node{}

		for len(temp) > 0 {
			n := temp[0]
			for i := 0; i < len(n.Children); i++ {
				s = append(s, n.Children[i])
			}
			temp = temp[1:]
		}
		ans++
	}

	return ans
}

func bfs2(root *Node) int {
	if root == nil {
		return 0
	}
	s := make([]*Node, 0)
	s = append(s, root)
	ans := 0
	for len(s) > 0 {
		length := len(s)
		for i := 0; i < length; i++ {
			n := s[0]
			for i := 0; i < len(n.Children); i++ {
				s = append(s, n.Children[i])
			}
			s = s[1:]
		}
		ans++
	}
	return ans
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	l := depth(node.Left) + 1
	r := depth(node.Right) + 1

	if l > r {
		return l
	} else {
		return r
	}
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findDuplicateSubtrees(root *TreeNode) (ans []*TreeNode) {
	m := make(map[string]*TreeNode)

	var dfs func(*TreeNode) string
	dfs = func(node *TreeNode) string {
		if node == nil {
			return ""
		}
		key := ""
		key += strconv.Itoa(node.Val)
		key += "/"
		key += dfs(node.Left)
		key += "/"
		key += dfs(node.Right)
		key += "/"

		if _, ok := m[key]; ok {
			ans = append(ans, node)
		} else {
			m[key] = node
		}
		return key
	}

	dfs(root)

	return
}
