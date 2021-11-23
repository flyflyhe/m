package dfs

import (
	"fmt"
	"sort"
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
