package heap

import "fmt"

func SortByMaxHeap(nums []int) []int{
	heap := BuildMaxHeap(nums)
	l := len(nums)
	for i := 0; i < l; i++{
		nums[l - i -1] = heap[0]
		heap[0], heap[l - i - 1] = heap[l - i - 1], heap[0]
		heap = BuildMaxHeap(heap[0:l - i - 1])
	}

	return nums
}

func BuildMaxHeap(nums []int) []int {
	heapArr := make([]int, len(nums) + 1)

	for i, num := range nums {
		 InsertHeap(heapArr, i + 1, num)
	}

	return heapArr[1:]
}

func InsertHeap(heapArr []int, i, num int) {
	if i == 1 {
		heapArr[1] = num
	} else {
		parent := i / 2
		if heapArr[parent] > num {
			heapArr[i] = num
		} else {
			heapArr[i] = heapArr[parent]
			heapArr[parent] = num
			InsertHeap(heapArr, parent, num)
		}
	}
}

func SortByMaxHeap2(nums []int) []int{
	BuildMaxHeap2(nums, len(nums))
	l := len(nums)
	for i := 0; i < l; i++ {
		l = l - 1
		nums[l], nums[i] = nums[i], nums[l]
		fmt.Println(nums)
		BuildMaxHeap2(nums, l)
		i--
	}

	return nums
}

func BuildMaxHeap2(nums []int, length int) {
	for i, _ := range nums {
		if i < length {
			InsertHeap2(nums, i)
		}
	}
}

func InsertHeap2(nums []int, i int) {
	if i == 0 {
		return
	}

	var pos int
	if i % 2 == 1 {
		pos = i / 2
	} else {
		pos = i / 2 - 1
	}

	if nums[pos] < nums[i] {
		nums[pos], nums[i] = nums[i], nums[pos]
		InsertHeap2(nums, pos)
	}
}
