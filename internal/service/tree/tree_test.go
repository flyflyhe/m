package tree

import (
	"fmt"
	"testing"
)

var tree *TreeNode
var tree2 *TreeNode

func init()  {
	tree = &TreeNode{Val: 10}
	tree.Left = &TreeNode{Val: 8, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 6}}
	tree.Right = &TreeNode{Val: 12, Right: &TreeNode{Val: 13}}

	tree2 = &TreeNode{Val: 1}
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