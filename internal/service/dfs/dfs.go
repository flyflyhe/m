package dfs

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func Dfs(startVertex int, graph [][]int)  {

	fmt.Println(graph)
	graphLen := len(graph)
	vertexVisited := make([]int, graphLen)
	for i := 0; i < graphLen; i++ {
		vertexVisited[i] = -1
	}

	var dfs func(int)
	dfs = func(startVertex int) {
		if startVertex > graphLen - 1 {
			return
		}

		for k, v := range vertexVisited {
			if startVertex == k && v == 1 {
				return
			}
		}

		vertexVisited[startVertex] = 1

		for i := 0; i < graphLen; i++ {
			if graph[startVertex][i] == 1 {
				fmt.Println("vertex", startVertex, "edge", i)
				dfs(i)
			}
		}
	}

	dfs(startVertex)
}

func CreateGraph(vertexNum int, relation [][]int) [][]int  {
	graph := make([][]int, vertexNum)
	for i := 0; i < vertexNum; i++ {
		graph[i] = make([]int, vertexNum)
	}

	//邻接矩阵
	for i := 0; i < len(relation); i++ {
		graph[relation[i][0]][relation[i][1]] = 1
	}

	return graph
}

/**
三数之和 回溯超时
 */

func threeSum(nums []int) (ret [][]int) {
	m := make(map[[3]int]bool)
	var dfs func([]int, []int, int)
	dfs = func(nums []int, path []int, start int) {
		if len(path) == 3 {
			tmp := 0
			tmpArr := [3]int{}
			for _, v := range path {
				tmp += v
			}
			if tmp == 0 {
				sort.Ints(path)
				for i, v := range path {
					tmpArr[i] = v
				}
				m[tmpArr] = true
			}
		}

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(nums, path, i+1)
			path = path[:len(path)-1]
		}
	}

	dfs(nums, []int{}, 0)

	for tmpArr, _ := range m{
		ret = append(ret, []int{tmpArr[0], tmpArr[1], tmpArr[2]})
	}

	return
}

/**
131:分割回文串
 */

func partition(s string) [][]string {
	var ret [][]string
	var dfs func(string, []string, int, *[][]string)
	dfs = func(s string, path []string, start int, ret *[][]string) {
		if start == len(s) {
			//进行一次切片拷贝，怕之后的操作影响tmpString切片内的值
			t := make([]string, len(path))
			copy(t, path)
			*ret = append(*ret, t)
		}

		for i := start; i < len(s); i++ {
			if isPartition(s, start, i) {
				path = append(path, s[start:i+1])
			} else {
				continue
			}

			dfs(s, path, i+1, ret)
			path = path[:len(path)-1]//因为i+1所以要弹出 i-n的字符串 防止重复
		}
	}
	dfs(s, []string{}, 0, &ret)

	return ret
}

//判断是否为回文
func isPartition(s string,startIndex,end int)bool{
	left:=startIndex
	right:=end
	for ;left<right;{
		if s[left]!=s[right]{
			return false
		}
		//移动左右指针
		left++
		right--
	}
	return true
}

func FindLHS(nums []int) int {
	var dfs func([]int, int, []int, bool, bool)
	ans := 0
	length := len(nums)

	dfs = func(nums []int, start int, path []int, up bool, check bool) {
		if start == length {
			if len(path) > 1 {
				if check {
					c := false
					for i := 1; i < len(path); i++ {
						if path[i] != path[0] {
							c = true
						}
					}

					if c {
						fmt.Println(path)
						ans = max(ans, len(path))
					}
				} else {
					ans = max(ans, len(path))
				}

			}
			return
		}

		for i := start; i < length; i++ {
			if len(path) == 0 {
				path = append(path, nums[i])
			} else {
				if up && nums[i] - path[0] == 1 {
					path = append(path, nums[i])
					check = false
				} else if !up &&  nums[i] - path[0] == -1 {
					path = append(path, nums[i])
					check = false
				} else if nums[i] - path[0] == 0 {
					path = append(path, nums[i])
				}
			}
			dfs(nums, i+1, path, up, check)
			path = []int{}
		}
	}

	dfs(nums, 0, []int{}, false, true)
	dfs(nums, 0, []int{}, true, true)

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func findLHS2(nums []int) int {
	sort.Ints(nums)
	fmt.Println(nums)
	ans := 0
	temp := 0
	effective := false
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[j] - nums[i] == 1  {
				temp++
				effective = true
			} else if nums[j] - nums[i] == 0 {
				temp++
			}
		}
		if effective {
			ans = max(ans, temp)
		}
		temp = 0
		effective = false
	}

	return ans
}

/*
:306 累加数
 */

func IsAdditiveNumber(num string) bool {
	var dfs func(sep int , num string) bool

	dfs = func(sep int, num string) bool {
		for i := sep; i < len(num) - 1; i++ {
			for step := 1; step < len(num) - sep; step++ {
				start, _ := strconv.Atoi(num[:sep+1])
				next, _ := strconv.Atoi(num[sep+1:sep+step+1])
				if len(strconv.Itoa(start)) != sep+1 {
					continue
				}
				if len(strconv.Itoa(next)) != step {
					continue
				}
				for j := sep+step+1; j < len(num);  {
					equal := start + next
					fmt.Println("start", start, "next", next, "equal", equal, "j", j, "equal", equal)

					if j + len(strconv.Itoa(equal)) <= len(num) {
						if strconv.Itoa(equal) == num[j:j + len(strconv.Itoa(equal))] {
							if j + len(strconv.Itoa(equal)) == len(num) {
								return true
							}
							j += len(strconv.Itoa(equal))
							fmt.Println("nextJ", j)
							start = next
							next = equal
							continue
						}
					}

					fmt.Println("break")
					break
				}
			}

			return dfs(i+1, num)
		}

		return false
	}

	return dfs(0, num)
}

/**
334 递增的三元子序列
 */

func IncreasingTriplet(nums []int) bool {
	var dfs func(start int, path []int, nums []int) bool
	dfs = func(start int, path []int, nums []int) bool {
		for i := start; i < len(nums); i++ {
			for j := start; j < len(nums); j++ {
				if len(path) == 0 {
					path = append(path, nums[j])
				} else if len(path) == 1 {
					if nums[j] > path[0] {
						path = append(path, nums[j])
					}
				} else if len(path) == 2 {
					if nums[j] > path[1] {
						path = append(path, nums[j])
						break
					} else if nums[j] > path[0] {
						if dfs(j + 1, []int{path[0], nums[j]}, nums) {
							return true
						}
					}
				}
			}

			if len(path) == 3 {
				return true
			}
			return dfs(i+1, []int{}, nums)
		}

		return false
	}

	return dfs(0, []int{}, nums)
}

func IncreasingTriplet2(nums []int) bool {
	stacks := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		needNewStack := true
		fmt.Println(stacks)
		for j := 0; j < len(stacks); j++ {
			if stacks[j][len(stacks[j])-1] < nums[i] {
				needNewStack = false
				if len(stacks[j]) == 2 {
					return true
				}
				stacks[j] = append(stacks[j], nums[i])

			} else if stacks[j][len(stacks[j])-1] == nums[i] {
				needNewStack = false
			} else if len(stacks[j]) > 1 &&  stacks[j][len(stacks[j])-2] < nums[i] {
				needNewStack = false
				stacks[j][1] = nums[i]
			}
		}

		if needNewStack {
			stacks = append(stacks, []int{nums[i]})
		}
	}

	return false
}

func IncreasingTriplet3(nums []int) bool {
	first := math.MaxInt32
	second := first
	for _, v :=  range nums {
		if v <= first {
			first = v
		} else if v <= second {
			second = v
		} else {
			return true
		}
	}
	return false
}