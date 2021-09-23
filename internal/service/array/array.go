package array

import "math"

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
		mid := (left + right + 1) / 2
		if arr[mid - 1] < arr[mid]  {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return right
}

func BinarySearch(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		mid := (l + r) / 2
		//log.Println("l", l, "r", r, "mid", mid)
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	//log.Println("l", l, "r", r)
	if nums[r] == target {
		return r
	} else {
		return -1
	}
}

func findMaxAverage(nums []int, k int) float64 {
	var ans float64
	var sum float64

	for i := 0; i < k; i++ {
		sum += float64(nums[i])
	}

	ans = sum / float64(k)

	for i := k; i < len(nums); i++ {
		sum = sum + float64(nums[i] - nums[i-k])
		ans = math.Max(ans, sum/float64(k))
	}

	return ans
}

func mergeSortArr(nums1 []int, m int, nums2 []int, n int)  {
	i := m - 1
	j := n - 1
	idx := m + n - 1
	for i >= 0 || j >= 0 {
		if i >= 0 && j >= 0 {
			if nums1[i] > nums2[j] {
				nums1[idx] = nums1[i]
				i--
			} else {
				nums1[idx] = nums2[j]
				j--
			}
		} else if i >= 0 {
			nums1[idx] = nums1[i]
			i--
		} else {
			nums1[idx] = nums2[j]
			j--
		}
		idx--
	}
}