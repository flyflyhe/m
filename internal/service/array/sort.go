package array

import (
	"fmt"
	"sort"
)

var tmp []int

func SortArray(nums []int) []int {
	tmp = make([]int, len(nums))
	mergeSort(nums, 0, len(nums))
	return nums
}

//归并排序

func mergeSort(nums []int, l, r int) {
	if l >= (r - 1) {
		return
	}

	mid := (l + r - 1) >> 1
	fmt.Println("l", l, "r", r, "m", mid)
	mergeSort(nums, l, mid+1)
	mergeSort(nums, mid+1, r)
	i := l
	j := mid + 1
	cnt := 0
	for ; i < mid+1 && j < r; cnt++ {
		if nums[i] <= nums[j] {
			tmp[cnt] = nums[i]
			i++
		} else {
			tmp[cnt] = nums[j]
			j++
		}
	}
	for ; i < mid+1; cnt++ {
		tmp[cnt] = nums[i]
		i++
	}
	for ; j < r; cnt++ {
		tmp[cnt] = nums[j]
		j++
	}

	for k := 0; k < r-l; k++ {
		nums[k+l] = tmp[k]
	}
}

func MergeSort(nums []int) []int {
	dst := make([]int, len(nums))
	var sort func([]int, []int, int, int)
	sort = func(nums, dst []int, start int, end int) {
		if start+1 >= end {
			return
		}
		mid := (start + end) / 2
		sort(dst, nums, start, mid)
		sort(dst, nums, mid, end)
		i := start
		j := mid
		k := start
		for i < mid || j < end {
			if j == end || (i < mid && nums[i] < nums[j]) {
				dst[k] = nums[i]
				i++
			} else {
				dst[k] = nums[j]
				j++
			}
			k++
		}
		fmt.Println(dst)
	}
	copy(dst, nums)
	sort(nums, dst, 0, len(nums))
	copy(nums, dst)
	return dst
}

func MergeSort2(nums []int) {
	length := len(nums)
	dst := make([]int, length)

	for seg := 1; seg < length; seg += seg {
		for start := 0; start < length; start += seg * 2 {
			end := min(start+seg*2, length)
			mid := min(start+seg, length)
			fmt.Println("seg", seg, "start", start, "mid", mid, "end", end)
			i := start
			j := mid
			k := start
			for i < mid || j < end {
				if j == end || (i < mid && nums[i] < nums[j]) {
					dst[k] = nums[i]
					i++
				} else {
					dst[k] = nums[j]
					j++
				}
				k++
			}
			fmt.Println(dst)
		}
		for i := 0; i < length; i++ {
			nums[i] = dst[i]
		}
	}
	copy(nums, dst)
}

func minimumAbsDifference(arr []int) [][]int {
	if len(arr) < 2 {
		return nil
	}

	sort.Ints(arr)

	m := make(map[int]int)

	min := arr[1] - arr[0]

	last := arr[1]
	var tmp int
	m[arr[0]] = arr[0]
	m[arr[1]] = arr[1]
	for i := 2; i < len(arr); i++ {
		m[arr[i]] = arr[i]
		tmp = arr[i] - last
		if tmp < min {
			min = tmp
		}
		last = arr[i]
	}

	result := make([][]int, 0)
	for _, v := range arr {
		if t, ok := m[v+min]; ok {
			result = append(result, []int{v, t})
		}
	}

	return result
}
