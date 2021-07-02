package dfs

import "fmt"

func Done() {
	var relation [][]int
	relation = append(relation, []int{0, 2})
	relation = append(relation, []int{2, 1})
	relation = append(relation, []int{3, 4})
	relation = append(relation, []int{2, 3})
	relation = append(relation, []int{1, 4})
	relation = append(relation, []int{2, 0})
	relation = append(relation, []int{0, 4})

	NumWays(5, relation, 3)
}

//{[[0,2],[2,1],[3,4],[2,3],[1,4],[2,0],[0,4]]}
func NumWays(n int, relation [][]int, k int) int {
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
	}

	for i := 0; i < len(relation); i++ {
		graph[relation[i][0]][relation[i][1]] = 1
	}

	fmt.Println(graph)

	return  1
}
