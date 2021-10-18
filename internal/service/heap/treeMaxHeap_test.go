package heap

import (
	"fmt"
	"testing"
)

func TestMaxHeap(t *testing.T)  {
	maxHeap := MaxHeap{}
	maxHeap.Insert(1)
	maxHeap.Insert(2)
	fmt.Println(PreOrder(maxHeap.GetRoot()))
	maxHeap.Insert(3)
	maxHeap.Insert(4)
	fmt.Println(PreOrder(maxHeap.GetRoot()))

	fmt.Println(maxHeap.Remove())
	fmt.Println(PreOrder(maxHeap.GetRoot()))
	fmt.Println(maxHeap.Remove())
	fmt.Println(PreOrder(maxHeap.GetRoot()))
	fmt.Println(maxHeap.Remove())
	fmt.Println(PreOrder(maxHeap.GetRoot()))
	fmt.Println(maxHeap.Remove())
	fmt.Println(PreOrder(maxHeap.GetRoot()))
}
