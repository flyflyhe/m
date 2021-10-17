package tree

import (
	"fmt"
	"testing"
)

var tree *TreeNode
var tree2 *TreeNode
var tree3 *TreeNode

func init()  {
	tree = &TreeNode{Val: 10}
	tree.Left = &TreeNode{Val: 8, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 6}}
	tree.Right = &TreeNode{Val: 12, Right: &TreeNode{Val: 13}}

	tree2 = &TreeNode{Val: 1}

	tree3 = &TreeNode{Val: 3}
}

func TestQueueMiddleTree(t *testing.T) {
	QueueMiddleTree(tree)
}

func TestInorderTraversal(t *testing.T) {
	fmt.Println(InorderTraversal(tree))
}

func TestTreeHeight(t *testing.T) {
	fmt.Println(TreeHeight(tree2))
	fmt.Println(TreeHeight(tree))
	fmt.Println(TreeHeight2(tree2))
	fmt.Println(TreeHeight2(tree))
}

func TestIsBalanced(t *testing.T) {
	fmt.Println(IsBalanced(tree))
	fmt.Println(IsBalanced(tree2))
}

func TestBuildTree(t *testing.T) {
	root := BuildTree([]int{3,9,20,15,7}, []int{9,3,15,20,7})
	fmt.Println(PreOrder(root))
}

func TestGetK(t *testing.T) {
	fmt.Println(GetK([]int{1,2,3,4}, 1))
}

func TestKthSmallest(t *testing.T) {
	fmt.Println(KthSmallest(tree, 3))
}