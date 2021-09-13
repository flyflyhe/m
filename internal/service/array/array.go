package array

func peakIndexInMountainArray(arr []int) int {
	point := 0
	for i := 1; i < len(arr); i++ {
		if arr[i-1] < arr[i] {
			point = i
		}
	}

	return point
}


func peakIndexInMountainArray2(arr []int) int {
	left := 1
	right := len(arr) - 1

	for left < right {
		mid := left + right + 1 >> 1
		if arr[mid - 1] < arr[mid]  {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return right
}