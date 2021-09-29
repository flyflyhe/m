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

/**
跳跃游戏
 */
func jump(nums []int) int {
	end := 0
	maxPosition := 0
	step := 0

	for i := 0; i < len(nums) - 1; i++ {
		maxPosition = max(maxPosition, i + nums[i])
		if i == end {
			step++
			end = maxPosition
		}
	}

	return step
}

/**
f(n) = f(n - k) + 1
 */
func jumpDp(nums []int) int  {
	arr := make([]int, len(nums))
	j := 0
	for i := 1; i < len(nums); i++ {
		for j + nums[j] < i {
			j++
		}
		arr[i] = arr[j] + 1
	}

	return arr[len(nums)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

/**
杨辉三角
 */
func GetRow(rowIndex int) []int {
	arr := make([][]int, rowIndex + 1)
	arr[0] = make([]int, 1)
	arr[0][0] = 1
	for i := 1; i < rowIndex + 1; i++ {
		arr[i] = make([]int, i + 1)
		arr[i][0] = 1
		arr[i][i] = 1
		for j := 1; j < i; j++ {
			arr[i][j] = arr[i-1][j-1] + arr[i-1][j]
 		}
	}

	return arr[rowIndex]
}

/**
超出时间限制
 */
func maxArea(height []int) int {
	maxArea := 0
	length := len(height)
	for i := 0; i < length - 1; i++ {
		for j := i + 1; j < length; j++ {
			maxArea = max(maxArea, (j - i) * min(height[i], height[j]))
		}
	}

	return maxArea
}

func maxArea2(height []int) int  {
	maxArea := 0
	length := len(height)
	l := 0
	r := length - 1
	for l < r {
		maxArea = max(maxArea, (r - l) * min(height[l], height[r]))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxArea
}
