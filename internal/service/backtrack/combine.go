package backtrack

var ret [][]int

func Combine(n int, k int) [][]int {
	ret = nil //开始时清空
	backtracking(n, k, 1, []int{})
	return ret
}

func backtracking(n, k, startIndex int, combination []int)  {
	if len(combination) == k {
		var b []int //注意使用深拷贝
		b = append(b, combination...)
		ret = append(ret, b)
		return
	}

	for i := startIndex; i <= n; i++ { // 控制树的横向遍历
		combination = append(combination, i) // 处理节点
		backtracking(n, k, i + 1, combination) // 递归：控制树的纵向遍历，注意下一层搜索要从i+1开始
		combination = combination[0:len(combination)-1] // 回溯，撤销处理的节点
	}
}
