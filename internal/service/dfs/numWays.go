package dfs

import "fmt"

func NumWays(n int, relation [][]int, k int) int {
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
	}

	//邻接矩阵
	for i := 0; i < len(relation); i++ {
		graph[relation[i][0]][relation[i][1]] = 1
	}

	fmt.Println(graph)

	var ways int

	var dfs func(int, int)
	dfs = func(startVertex, step int) {
		if step == k {
			if startVertex == n - 1{
				ways++
			}
			return
		}
		fmt.Println("v", startVertex, "s", step)
		for i := 0; i < len(graph); i++ {
			to := graph[startVertex][i]
			if to == 1 {
				dfs(i, step + 1)
			}
		}
	}

	dfs(0, 0)

	return ways
}
