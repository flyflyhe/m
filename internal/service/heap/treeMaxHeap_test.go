package heap

import (
	"fmt"
	"testing"
)

func TestMaxHeap(t *testing.T)  {
	maxHeap := MaxHeap{}
	maxHeap.Insert(1, 11)
	maxHeap.Insert(2, 22)
	fmt.Println(PreOrder(maxHeap.GetRoot()))
	maxHeap.Insert(3, 33)
	maxHeap.Insert(4, 44)
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

func TestTopKFrequent(t *testing.T) {
	fmt.Println(TopKFrequent([]int{1,1,1,2,2,2,3,3,3}, 3))
}
