package array

import (
	"fmt"
	"log"
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

type IntSlice [][]int

func (x IntSlice) Len() int           { return len(x) }
func (x IntSlice) Less(i, j int) bool { return x[i][0] < x[j][0] }
func (x IntSlice) Swap(i, j int)      { x[i][0], x[j][0] = x[j][0], x[i][0] }

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	p := make(IntSlice, 0)
	p = append(p, p1, p2, p3, p4)
	sort.Sort(p)

	return abc(p[0][0], p[1][0]) == abc(p[2][0], p[3][0]) || abc(p[0][0], p[3][0]) == abc(p[1][1], p[2][1])
}

func abc(a, b int) int {
	if a > 0 && b < 0 {
		return a - b
	} else if a < 0 && b > 0 {
		return b - a
	} else {
		t := a - b
		if t < 0 {
			return -t
		} else {
			return t
		}
	}
}

func minSubsequence(nums []int) []int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	var dfs func([]int, []int)
	var s [][]int
	dfs = func(nums []int, tmpS []int) {
		for i := 0; i < len(nums); i++ {
			tmpS = append(tmpS, nums[i])
			var t []int
			t = append(t, tmpS...)
			s = append(s, t)
			for j := i + 1; j < len(nums); j++ {
				tmpS = append(tmpS, nums[j])
				var t []int
				t = append(t, tmpS...)
				s = append(s, t)
				if j+1 < len(nums) {
					dfs(nums[j+1:], tmpS)
				}
				tmpS = tmpS[:len(tmpS)-1]
			}
			tmpS = tmpS[:len(tmpS)-1]
		}
	}

	dfs(nums, []int{})
	log.Println(s)
	log.Println(len(s))

	return nil
}

func Dfs(temp, nums []int, start int) {
	tmp := make([]int, len(temp))
	copy(tmp, temp)
	//res = append(res, tmp)
	for i := start; i < len(nums); i++ {
		//if i>start&&nums[i]==nums[i-1]{
		//	continue
		//}
		temp = append(temp, nums[i])
		Dfs(temp, nums, i+1)
		temp = temp[:len(temp)-1]
	}
}
