package array

import (
	"fmt"
	"math"
)

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

func destCity(paths [][]string) string {
	return ""
}

func ThirdMax(nums []int) int {
	max := -1<<31
	second := -1<<31
	third := -1<<31
	hasThird := false

	for _, v := range nums {
		if max < v {
			max = v
		}
	}
	for _, v := range nums {
		if  v < max && second < v {
			second = v
		}
	}
	for _, v := range nums {
		if v < second {
			hasThird = true
			if third < v {
				third = v
			}
		}
	}

	if hasThird {
		return third
	}

	return max
}

func ThirdMax2(nums []int) int {
	a := -1<<32
	b := -1<<32
	c := -1<<32

	for _, x := range nums {
		if x > a {
			c = b; b = a; a = x
		} else if x < a && x > b {
			c = b; b = x
		} else if x < b && x > c {
			c = x
		}
	}

	if c != -1 << 32 {
		return c
	}

	return a
}

/**
山脉数组
 */

func PeakIndexInMountainArray(arr []int) int {
	length := len(arr)
	l := 1
	r := length - 2
	for l < r {
		mid := (l + r) / 2
		fmt.Println("l", l, "r", r, "m",mid)
		if arr[mid-1] < arr[mid] && arr[mid] < arr[mid+1] {
			l = mid + 1
		} else if arr[mid - 1] > arr[mid] && arr[mid] > arr[mid + 1] {
			r = mid
		} else {
			return mid
		}
	}

	return r
}

/**
二维数组的遍历
 */

func TwoArrayErgodic(s [][]int)  {
	rows := len(s)
	columns := len(s[0])
	totalElements := rows * columns
	order :=  make([]int, totalElements)
	vertical := -1
	horizontal := 1
	row := 0
	column := 0
	for i := 0; i < totalElements; i++ {
		order[i] = s[row][column]
		newRow := row + vertical
		newColumn := column + horizontal
		if newColumn >= columns {
			row++
			vertical = 1
			horizontal = -1
		} else if newRow < 0 {
			column++
			vertical = 1
			horizontal = -1
		} else if newRow >= rows {
			column++
			vertical = -1
			horizontal = 1
		} else if newColumn < 0 {
			row++
			vertical = -1
			horizontal = 1
		} else {
			row = newRow
			column = newColumn
		}
	}
	fmt.Println(order)
}

func TwoArrayErgodic2(s [][]int)  {
	rows := len(s)
	columns := len(s[0])
	totalElements := rows * columns
	order :=  make([]int, totalElements)
	vertical := -1 // 竖直
	horizontal := -1 //水平
	row := 0
	column := columns - 1
	for i := 0; i < totalElements; i++ {
		order[i] = s[row][column]
		newRow := row + vertical
		newColumn := column + horizontal
		if newRow < 0 {
			column--
			vertical = 1
			horizontal = 1
		} else if newColumn >= columns {
			row++
			vertical = -1
			horizontal = -1
		} else if newRow >= rows {
			column--
			vertical = -1
			horizontal = -1
		} else if newColumn < 0 {
			row++
			vertical = 1
			horizontal = 1
		} else {
			row = newRow
			column = newColumn
		}
	}
	fmt.Println(order)
}

func ShoppingOffers(price []int, special [][]int, needs []int) int {
	//1.计算原价
	//2.计算大礼包
	//3.更具最低价个减去需求单后，递归计算购物单
	return dfs(price,special,needs)

}
func dfs(prices []int,specials [][]int,needs []int) int{
	minPrice:=0
	//计算原价
	for i,v:=range needs{
		minPrice +=prices[i]*v
	}
	//计算大礼包
loop:
	for _,s:=range specials{
		newNeeds:=make([]int,len(needs))
		copy(newNeeds, needs)
		/**
		每一个礼包 用一个用两个用n个循环
		 */
		for i:=range newNeeds{
			newNeeds[i]-=s[i]
			if newNeeds[i]<0{
				continue loop
			}
		}
		curPrice:=dfs(prices,specials,newNeeds)+s[len(prices)]
		minPrice =minInt(minPrice,curPrice)
	}
	return minPrice

}

func minInt(values ...int)int{
	result:=math.MaxInt32
	for _,v:=range values{
		if v<result{
			result = v
		}
	}
	return result
}

func Combine(n, k int) [][]int {
	var ret [][]int
	var c []int
	combineDfs(0, n, k, c, &ret)
	return ret
}

/**
	void backtracking(参数) {
		if (终止条件) {
			存放结果;
			return;
		}

		for (选择：本层集合中元素（树中节点孩子的数量就是集合的大小）) {
			处理节点;
			backtracking(路径，选择列表); // 递归
			回溯，撤销处理结果
		}
	}
 */

func combineDfs(start, n, k int, c []int, ret *[][]int) {
	if len(c) == k {
		tmp := make([]int, len(c))
		copy(tmp, c)
		*ret = append(*ret, tmp)
		fmt.Println(tmp)
		return
	}
	for i := start; i < n - (k - len(c)) + 1 ; i++ {
		c = append(c, i)
		combineDfs(i + 1, n, k, c, ret)
		c = c[0:len(c) - 1]
	}
}

func CombinationSum3(k int, n int) [][]int {
	var ret [][]int
	var path []int
	combinationSumDfs(1, 10, n, k, path, &ret)
	return ret
}

func combinationSumDfs(start, end, n , k int, path []int, ret *[][]int) {
	if len(path) == k && Sum(path...) == n {
		tmp := make([]int, k)
		copy(tmp, path)
		*ret = append(*ret, tmp)
		return
	}

	for i := start; i < end ; i++ {
		path = append(path, i)
		combinationSumDfs(i+1, end, n, k, path, ret)
		path = path[0:len(path)-1]
	}
}

func Sum(values... int) (s int) {
	for _, item := range values {
		s += item
	}
	return
}

func GenerateMatrix(n int) [][]int {
	cols := n
	rows := n
	matrix := make([][]int, n)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, n)
	}
	row := 0
	col := -1
	moveArr := [][]int {
		{0, 1}, //向右
		{1, 0}, //向下
		{0, -1}, //向左
		{-1, 0}, //向上
	}
	direction := 0
	for i := 1; i <= cols * rows; i++ {
		row = row + moveArr[direction][0]
		col = col + moveArr[direction][1]

		//fmt.Println("row",row, "col", col)
		nextRow := row + moveArr[direction][0]
		nextCol := col + moveArr[direction][1]
		if nextRow == cols  {
			direction++
		} else if nextCol == rows {
			direction++
		} else if nextCol < 0 {
			direction++
		} else if matrix[nextRow][nextCol] != 0 {
			direction++
		}

		direction = direction & 3
		matrix[row][col] = i
	}

	return matrix
}

func MaximalRectangle(matrix []string) int {
	area := 0
	debug := true
	rows := len(matrix)
	if rows == 0 {
		return 0
	}
	cols := len(matrix[0])
	fmt.Println("rows", rows, "cols", cols)
	rowEnd := rows - 1
	colEnd := cols - 1
	/**
	起点 {x,y} 遍历所有比它大的坐标
	 */
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if colEnd > rowEnd {
				for col1 := colEnd; col1 >= col; col1-- {
					for row1 := rowEnd; row1 >= row; row1-- {
						tmp, ok := MatrixArea(row, col, row1, col1, matrix)
						if ok {
							if debug {
								fmt.Println([]int{row, col, row1, col1})
								fmt.Println(tmp)
							}

							if tmp > area {
								area = tmp
							}
						} else {
							if tmp <= area {
								goto next1
							}
						}
					}
					next1:
				}
			} else {
				for row1 := rowEnd; row1 >= row; row1-- {
					for col1 := colEnd; col1 >= col; col1-- {
						tmp, ok := MatrixArea(row, col, row1, col1, matrix)
						if ok {
							if debug {
								fmt.Println([]int{row, col, row1, col1})
								fmt.Println(tmp)
							}
							if tmp > area {
								area = tmp
							}
						} else {
							if tmp <= area {
								goto next2
							}
						}
					}
					next2:
				}
			}
		}
	}

	return area
}

func MatrixArea(x, y, x1, y1 int, matrix []string) (int, bool) {
	area := (x1 - x + 1) * (y1 - y + 1)
	result := true
	for row := x ; row <= x1; row++ {
		for col := y ; col <= y1; col++ {
			if matrix[row][col] == '0' {
				result = false
				break
			}
		}
	}
	return area, result
}

/**
快速排序
 */

func KSort(nums []int)  {
	var sort func(int, int)

	sort = func(left, right int) {
		if left > right {
			return
		}
		tmp := nums[left]
		l := left
		r := right
		for l < r {
			for l < r && nums[r] >= tmp {
				r--
			}
			for l < r && nums[l] <= tmp {
				l++
			}
			if l < r {
				nums[l], nums[r] = nums[r], nums[l]
			}
		}
		nums[left] = nums[l]
		nums[l] = tmp
		sort(left, l - 1)
		sort(l + 1, right)
	}

	sort(0, len(nums) - 1)
}

func KSort2(nums []int)  {
	var sort func(int, int)
	sort = func(start, end int) {
		if end > start {
			random := start
			nums[random], nums[end] = nums[end], nums[random] //把选中的值放入尾部

			p1 := start - 1
			for i := start; i < end; i++ {
				if nums[i] < nums[end] {
					p1++
					nums[p1], nums[i]  = nums[i], nums[p1]
				}
			}
			//循环结束p1 执行小于random的最后一个值 所以需要++ p1 可能==0 没有值比random 小
			p1++
			nums[p1], nums[end] = nums[end], nums[p1]

			sort(start, p1 - 1)
			sort(p1 + 1, end)
		}
	}
	sort(0, len(nums) - 1)
}

/**
只出现一次的数字
 */

func singleNumber(nums []int) []int {
	xorSum := 0
	for _, num := range nums {
		xorSum ^= num
	}
	lsb := xorSum & -xorSum
	type1, type2 := 0, 0
	for _, num := range nums {
		if num&lsb > 0 {
			type1 ^= num
		} else {
			type2 ^= num
		}
	}
	return []int{type1, type2}
}