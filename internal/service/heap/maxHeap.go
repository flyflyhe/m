package heap

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
