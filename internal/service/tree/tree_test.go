package tree

import (
	"fmt"
	"log"
	"testing"
)

var tree *TreeNode
var tree2 *TreeNode
var tree3 *TreeNode
var tree4 *TreeNode

func init() {
	tree = &TreeNode{Val: 10}
	tree.Left = &TreeNode{Val: 8, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 6}}
	tree.Right = &TreeNode{Val: 12, Right: &TreeNode{Val: 13}}

	tree2 = &TreeNode{Val: 1}

	tree3 = &TreeNode{Val: 3}
	tree4 = &TreeNode{Val: 1, Left: &TreeNode{Val: 0, Left: &TreeNode{Val: 0}}, Right: &TreeNode{Val: 1}}
}

func Test_pruneTree(t *testing.T) {
	pruneTree(tree4)
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
	root := BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	fmt.Println(PreOrder(root))
}

func TestGetK(t *testing.T) {
	fmt.Println(GetK([]int{1, 2, 3, 4}, 1))
}

func TestKthSmallest(t *testing.T) {
	fmt.Println(KthSmallest(tree, 3))
}

func TestPostorderTraversalRecursion(t *testing.T) {
	fmt.Println(PostorderTraversalRecursion(tree))
	fmt.Println(PostorderTraversal(tree))
	fmt.Println(PreorderTraversal(tree))
}

func TestIsSymmetric(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 2}}}
	fmt.Println(IsSymmetric(root))
}

func Test_depth(t *testing.T) {
	log.Println(depth(tree))
	log.Println(depth(tree2))
}
