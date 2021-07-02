package redBlack

/**
红黑树特性
0:任然是bst
1:每个节点都有颜色 红色或黑色
2:跟节点是黑色的
3:每个叶子节点(NIL)都是黑色
4:如果一个节点是红色的，则它的子节点必须是黑色的
5:任意一个节点到该节点的每个叶子节点的所有路径上包含相同数目的黑节点; 此性质反应了红黑树的平衡性 它确保从任一个节点出发到其叶子节点的所有路径中，
其最长路径长度也不会超过最短路径长度的两倍.因而，红黑树是接近平衡的二叉树. 而且指出黑节点的层数是相等的 所以黑节点是完美平衡的
 */

const RedColor = "red"
const BlackColor = "black"

type RbNode struct {
	Color string
	Value int
	Left *RbNode
	Right *RbNode
	Parent *RbNode
}

func InitNode(value int, color string) *RbNode  {
	node := &RbNode{Value: value,Color: color}
	return node
}

type RbTree struct {
	MRoot *RbNode //根节点
}

// LeftRotate pNode支点(旋转节点)
func (rbTree *RbTree) LeftRotate(pNode *RbNode)  {
	vNode := pNode.Right
	rNode := vNode.Left

	pNode.Right = rNode
	if rNode != nil {
		rNode.Parent = pNode //修正R的parent 为P
	}
	vNode.Parent = pNode.Parent //修正v的parent 为p原来的的parent

	if pNode.Parent == nil { //如果pNode 没有父节点 说明pNode是根节点 替换根节点为vNode
		rbTree.MRoot = vNode
	} else { //如果pNode 有父节点 则v取代p成为这个父节点的左孩子或右孩子
		if pNode.Parent.Left == pNode {
			pNode.Parent.Left = vNode
		} else {
			pNode.Parent.Right = vNode
		}
	}

	vNode.Left = pNode
	pNode.Parent = vNode
}

func (rbTree *RbTree) RightRotate(pNode *RbNode) {
	fNode := pNode.Left
	kNode := fNode.Right

	pNode.Left = kNode
	if kNode != nil {
		kNode.Parent = pNode
	}

	fNode.Parent = pNode.Parent
	if pNode.Parent == nil {
		rbTree.MRoot = fNode
	} else {
		if pNode.Parent.Left == pNode {
			pNode.Parent.Left = fNode
		} else {
			pNode.Parent.Right = fNode
		}
	}

	fNode.Right = pNode
	pNode.Parent = fNode
}

