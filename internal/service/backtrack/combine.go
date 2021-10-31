package backtrack

import (
	"container/heap"
	"fmt"
	"sort"
)

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

func MaxTwoEvents(events [][]int) (ans int) {
	sort.Slice(events, func(i, j int) bool { return events[i][0] < events[j][0] }) // 按开始时间排序
	h := hp{}
	maxVal := 0
	for _, e := range events {
		start, end, val := e[0], e[1], e[2]
		for len(h) > 0 && h[0].end < start { // 如果结束时间早于当前活动开始时间
			maxVal = max(maxVal, heap.Pop(&h).(pair).val) // 更新前面可以选择的活动的最大价值
		} //因为end一定大于start 所以 之后的活动都可以加上 maxVal这个最大值
		fmt.Println("maxVal", maxVal, "val", val, "start", start, "end", end, "ans", ans)
		ans = max(ans, maxVal+val) // 至多参加两个活动
		heap.Push(&h, pair{end, val})
	}
	return
}

type pair struct{ end, val int }
type hp []pair
func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].end < h[j].end }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
